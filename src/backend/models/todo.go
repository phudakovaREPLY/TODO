package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Todo struct
type Todo struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
	Body  string             `json:"body" bson:"body"`
	Done  bool               `json:"done" bson:"done"`
}
