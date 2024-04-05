package kafka

import (
	"log"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	ConfigMap *ckafka.ConfigMap
	Topic    string
}

func NewKafkaProducer(configMap *ckafka.ConfigMap, topicName string) *KafkaProducer {
	return &KafkaProducer{
		ConfigMap: configMap,
		Topic:    topicName,
	}
}

func (c *KafkaProducer) Produce(serializedPayload []byte) error {
	producer, err := ckafka.NewProducer(c.ConfigMap)
	if err != nil {
		return err
	}
	err = producer.Produce(&ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &c.Topic, Partition: ckafka.PartitionAny},
		Value:          serializedPayload,
	}, nil)
	if err != nil {
		return err
	} else {
		log.Printf("Message successfully sent to topic %s", c.Topic)
	}
	return nil
}