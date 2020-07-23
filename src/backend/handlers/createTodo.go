package hanlders

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wa-brown/TODO/src/backend/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoClient to interact with mongoDB
var MongoClient *mongo.Client

// CreateTodo in the database
func CreateTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")


	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	collection := MongoClient.Database("todo").Collection("todos")

	result, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": Error - ` + err.Error() + `" }`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
