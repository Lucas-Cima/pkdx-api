package main

import (
	"context"
	"fmt"
	"log"
	"pkdx-api/pkg/controller"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDb *mongo.Collection

func main() {
	fmt.Println("SERVER UP")
	findOptions := options.Find()
	findOptions.SetLimit(1000)

	clientOptions := options.Client().ApplyURI("mongodb+srv://Lucas:Pokemon@pokedex.l4iml.mongodb.net/Pokedex?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	mongoDb = client.Database("Pokedex").Collection("Pokemon")
	controller.MongoDb = *mongoDb

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			logrus.Error(err)
		}
	}()
	controller.HandleRequests()
}
