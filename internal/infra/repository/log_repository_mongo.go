package repository

import (
	"context"
	"log"
	"encoding/json"
	"github.com/henriquemarlon/m9-p2/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
)

type LogRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewLogRepositoryMongo(client *mongo.Client, dbName string, collection string) *LogRepositoryMongo {
	sensorsColl := client.Database(dbName).Collection(collection)
	return &LogRepositoryMongo{
		Collection: sensorsColl,
	}
}

func (s *LogRepositoryMongo) CreateLog(sensorLog *entity.Log) error {
	result, err := s.Collection.InsertOne(context.TODO(), sensorLog)
	log.Printf("Inserting log into the MongoDB collection with id: %s", result)
	return err
}

func (s *LogRepositoryMongo) FindAllLogs() ([]*entity.Log, error) {
	cur, err := s.Collection.Find(context.TODO(), bson.D{})
	log.Printf("Selecting all Logs from the MongoDB collection %s", s.Collection.Name())
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var Logs []*entity.Log
	for cur.Next(context.TODO()) {
		var log bson.M
		err := cur.Decode(&log)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found")
		} else if err != nil {
			return nil, err
		}

		jsonLogData, err := json.MarshalIndent(log, "", " ")
		if err != nil {
			return nil, err
		}

		var logData entity.Log
		err = json.Unmarshal(jsonLogData, &logData)
		if err != nil {
			return nil, err
		}
		Logs = append(Logs, &logData)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return Logs, nil
}
