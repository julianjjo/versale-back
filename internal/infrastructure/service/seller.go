package service

import (
	"context"

	repository "github.com/julianjjo/versasale-back/internal/infrastructure/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Seller struct {
	SellerId   string `json:"seller_id" bson:"seller_id"`
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	Email      string `json:"email" bson:"email"`
	TypeId     string `json:"type_id" bson:"type_id"`
	DocumentId string `json:"document_id" bson:"document_id"`
	Age        int    `json:"age" bson:"age"`
	Password   string `json:"password" bson:"password"`
}

func SaveSeller(client *mongo.Client, ctx context.Context, seller Seller) error {
	err := repository.SaveToMongoDB(client, ctx, "seller", seller)
	if err != nil {
		return err
	}
	return nil
}
