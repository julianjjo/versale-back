package service

import (
	"context"

	model "github.com/julianjjo/versasale-back/internal/core/model"
	repository "github.com/julianjjo/versasale-back/internal/infrastructure/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveSeller(client *mongo.Client, ctx context.Context, seller model.Seller) error {
	err := repository.SaveToMongoDB(client, ctx, "Seller", seller)
	if err != nil {
		return err
	}
	return nil
}
