package db

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"pkdx-api/pkg/model"
)

// Function for querying the pokedex collection.
// Pulls the whole collection aka the National Dex and ALL Pokemon Forms
func GetPokedex(ctx context.Context, collection *mongo.Collection) (pokedex []model.Pokemon) {
	var pokemon model.Pokemon
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "Number", Value: 1}})
	res, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Error(err)
	}
	for res.Next(ctx) {
		err := res.Decode(&pokemon)
		if err != nil {
			log.Fatal(err)
		}
		pokedex = append(pokedex, pokemon)
	}
	return
}
