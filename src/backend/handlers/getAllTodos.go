package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wa-brown/TODO/src/backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllTodos from the database
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	var todos []models.Todo
	collection := MongoClient.Database("todo").Collection("todos")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var todo models.Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
