package service

import (
	"context"

	model "github.com/julianjjo/versasale-back/internal/core/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDocumentType(client *mongo.Client, ctx context.Context) ([]model.DocumentType, error) {
	collectionRef := client.Database("versasale").Collection("DocumentType")
	filter := bson.M{}

	var result []model.DocumentType
	cursor, err := collectionRef.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc model.DocumentType
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		result = append(result, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
