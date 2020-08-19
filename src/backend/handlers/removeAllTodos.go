package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wa-brown/TODO/src/backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateTodo in the database
func RemoveAllTodos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers, access-control-allow-methods")

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	collection := MongoClient.Database("todo").Collection("todos")

	result, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": Error - ` + err.Error() + `" }`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}