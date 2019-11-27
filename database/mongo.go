package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

var mongoClient *mongo.Client

func GetMongoClient() *mongo.Client {
	return mongoClient
}

func MongoConnect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client

	fmt.Println("Connected to MongoDB!")
}
