package repository

import (
	"context"
	"log"
	"github.com/henriquemarlon/m9-p2/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
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
