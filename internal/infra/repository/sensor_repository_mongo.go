package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/henriquemarlon/m9-p2/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type SensorRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewSensorRepositoryMongo(client *mongo.Client, dbName string, targetCollection string) *SensorRepositoryMongo {
	collection := client.Database(dbName).Collection(targetCollection)
	return &SensorRepositoryMongo{
		Collection: collection,
	}
}

func (s *SensorRepositoryMongo) FindAllSensors() ([]*entity.Sensor, error) {
	cur, err := s.Collection.Find(context.TODO(), bson.D{})
	log.Printf("Selecting all Sensors from the MongoDB collection %s", s.Collection.Name())
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var Sensors []*entity.Sensor
	for cur.Next(context.TODO()) {
		var sensor bson.M
		err := cur.Decode(&sensor)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found")
		} else if err != nil {
			return nil, err
		}

		jsonSensorData, err := json.MarshalIndent(sensor, "", " ")
		if err != nil {
			return nil, err
		}

		var sensorData entity.Sensor
		err = json.Unmarshal(jsonSensorData, &sensorData)
		if err != nil {
			return nil, err
		}
		Sensors = append(Sensors, &sensorData)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return Sensors, nil
}
