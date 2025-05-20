package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Author   string             `json:"author"`
	IsRented bool               `json:"is_rented,omitempty"`
}
