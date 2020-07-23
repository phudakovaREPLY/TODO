package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/wa-brown/TODO/src/backend/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	mongoClient, _ = mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err := mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Cannot connect to mongo-service: %+v", err)
	}

	fmt.Println("Connected to mongoDB")

	handlers.MongoClient = mongoClient

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/createtodo", handlers.CreateTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}
