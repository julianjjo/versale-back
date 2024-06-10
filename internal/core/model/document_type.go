package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DocumentType struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}
