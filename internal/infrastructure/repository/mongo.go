package repository

import (
	"context"
	"log/slog"
	"os"
	"time"

	config "github.com/julianjjo/versasale-back/internal/infrastructure/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectToMongoDB connects to MongoDB and returns the client, context, cancel function, and error if any
func ConnectToMongoDB() (*mongo.Client, context.Context, context.CancelFunc, error) {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDB.Connection))
	if err != nil {
		cancel()
		return nil, nil, nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		cancel()
		return nil, nil, nil, err
	}
	return client, ctx, cancel, nil
}

// SaveToMongoDB saves a customer to MongoDB
func SaveToMongoDB(client *mongo.Client, ctx context.Context, collectionName string, data interface{}) error {
	collection := client.Database("versasale").Collection(collectionName)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
