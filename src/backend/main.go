package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient to interact with mongoDB
var MongoClient *mongo.Client

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://mongo-service:27017")

	MongoClient, _ = mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err := MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Cannot connect to mongo-service: %+v", err)
	}

}
