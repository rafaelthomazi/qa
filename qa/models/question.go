package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Question represents a question made by one user
type Question struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Statement string             `json:"statement" bson:"statement"`
	Answer    string             `json:"answer" bson:"answer"`
}
