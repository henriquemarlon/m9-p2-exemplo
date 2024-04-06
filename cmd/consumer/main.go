package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/henriquemarlon/m9-p2/internal/infra/kafka"
	"github.com/henriquemarlon/m9-p2/internal/infra/repository"
	"github.com/henriquemarlon/m9-p2/internal/usecase"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	options := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s",
			os.Getenv("MONGODB_ATLAS_USERNAME"),
			os.Getenv("MONGODB_ATLAS_PASSWORD"),
			os.Getenv("MONGODB_ATLAS_CLUSTER_HOSTNAME"),
			os.Getenv("MONGODB_ATLAS_APP_NAME")))
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	msgChan := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"sasl.mechanisms":    "PLAIN",
		"security.protocol":  "SASL_SSL",
		"sasl.username":      os.Getenv("CONFLUENT_API_KEY"),
		"sasl.password":      os.Getenv("CONFLUENT_API_SECRET"),
		"session.timeout.ms": 6000,
		"group.id":           "m9-p2",
		"auto.offset.reset":  "latest",
	}

	kafkaRepository := kafka.NewKafkaConsumer(configMap, []string{os.Getenv("CONFLUENT_KAFKA_TOPIC_NAME")})
	go func() {
		if err := kafkaRepository.Consume(msgChan); err != nil {
			log.Printf("Error consuming kafka queue: %v", err)
		}
	}()

	logRepository := repository.NewLogRepositoryMongo(client, "mongodb", "sensors_log")
	createSensorLogUseCase := usecase.NewLogUseCase(logRepository)

	for msg := range msgChan {
		// fmt.Printf("Consumed message: %s\n", msg.Value)
		dto := usecase.CreateLogInputDTO{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		}
		err = createSensorLogUseCase.Execute(dto)
		if err != nil {
			log.Fatalf("Failed to create sensor log: %v", err)
		}
	}
}
