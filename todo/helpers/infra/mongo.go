package infra

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Db struct {
	MongoClient *mongo.Client
}

func ConnectMongo() *Db {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://puv0:HZ97zsG72mcxa@cluster0.533rj.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfuly Connected MongoDb")
	return &Db{MongoClient: client}
}

func (db *Db) GetCollection(collectionName string) *mongo.Collection {
	return db.MongoClient.Database("habitDB").Collection(collectionName)
}
