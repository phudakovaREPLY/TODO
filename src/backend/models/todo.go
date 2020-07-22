package models

// Todo struct
type Todo struct {
	Tite string `json:"title" bson:"title"`
	Body string `json:"body" bson:"body"`
}
