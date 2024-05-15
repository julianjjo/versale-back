package service

import (
	"context"
	"fmt"

	repository "github.com/julianjjo/versasale-back/internal/infrastructure/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Customer struct {
	CustomerId string `json:"customer_id" bson:"customer_id"`
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	Email      string `json:"email" bson:"email"`
	TypeId     string `json:"type_id" bson:"type_id"`
	DocumentId string `json:"document_id" bson:"document_id"`
	Age        int    `json:"age" bson:"age"`
	Password   string `json:"password" bson:"password"`
}

func SaveCustomer(client *mongo.Client, ctx context.Context, customer Customer) error {
	err := repository.SaveToMongoDB(client, ctx, "customer", customer)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func EmailExists(client *mongo.Client, ctx context.Context, email string, collection string) (bool, error) {
	collectionRef := client.Database("versasale").Collection(collection)
	filter := bson.M{"email": email}

	var result bson.M
	err := collectionRef.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func DocumentIdExists(client *mongo.Client, ctx context.Context, documentId string, collection string) (bool, error) {
	fmt.Println("DocumentIdExists()")
	collectionRef := client.Database("versasale").Collection(collection)
	filter := bson.M{"document_id": documentId}

	var result bson.M
	err := collectionRef.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("DocumentId does not exist")
			return false, nil
		}
		return false, err
	}
	fmt.Println(err)
	return true, nil
}
