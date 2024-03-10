package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSetup() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("failed to connect to MongoDB:", err)
		return nil
	}

	log.Println("Successfully connected to MongoDB")
	return client
}

var Client *mongo.Client = DBSetup()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	productCollection := client.Database("Ecommerce").Collection(collectionName)
	return productCollection
}
