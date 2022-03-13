package main

import (
	"context"
	"fmt"
	"log"
	"pkdx-api/pkg/routes"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	mongoDb = client.Database("Pokedex").Collection("Pokemon")
	routes.MongoDb = *mongoDb

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			logrus.Error(err)
		}
	}()
	routes.HandleRequests()
}
