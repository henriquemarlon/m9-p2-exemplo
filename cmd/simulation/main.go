package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"github.com/henriquemarlon/m9-p2/internal/domain/entity"
	"github.com/henriquemarlon/m9-p2/internal/usecase"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/henriquemarlon/m9-p2/internal/infra/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewSensorRepositoryMongo(client, "mongodb", "sensors")
	findAllSensorsUseCase := usecase.NewFindAllSensorsUseCase(repository)

	sensors, err := findAllSensorsUseCase.Execute()
	if err != nil {
		log.Fatalf("Failed to find all sensors: %v", err)
	}

	var wg sync.WaitGroup
	for _, sensor := range sensors {
		wg.Add(1)
		log.Printf("Starting sensor: %v", sensor)
		go func(sensor usecase.FindAllSensorsOutputDTO) {
			defer wg.Done()
			opts := MQTT.NewClientOptions().AddBroker(
				fmt.Sprintf("ssl://%s:%s",
					os.Getenv("BROKER_TLS_URL"),
					os.Getenv("BROKER_PORT"))).SetUsername(
				os.Getenv("BROKER_USERNAME")).SetPassword(
				os.Getenv("BROKER_PASSWORD")).SetClientID(sensor.ID)
			client := MQTT.NewClient(opts)
			if session := client.Connect(); session.Wait() && session.Error() != nil {
				log.Fatalf("Failed to connect to MQTT broker: %v", session.Error())
			}
			for {
				payload, err := entity.NewSensorPayload(
					sensor.ID,
					sensor.Unit,
					sensor.Params,
					time.Now(),
				)
				if err != nil {
					log.Fatalf("Failed to create payload: %v", err)
				}

				jsonBytesPayload, err := json.Marshal(payload)
				if err != nil {
					log.Println("Error converting to JSON:", err)
				}
 
				token := client.Publish(os.Getenv("BROKER_TOPIC"), 1, false, string(jsonBytesPayload))
				log.Printf("Published: %s, on topic: %s", string(jsonBytesPayload), os.Getenv("BROKER_TOPIC"))
				token.Wait()
				time.Sleep(10 * time.Second)
			}
		}(sensor)
	}
	wg.Wait()
}