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
func GetPokedex(ctx context.Context, collection *mongo.Collection) (pokedex []model.Pokemon, err error) {
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
	return pokedex, nil
}

func GetPokemonByDexNum(ctx context.Context, collection *mongo.Collection, dexNum string) (pokemonByDexNum []model.Pokemon, err error) {
	var pokemon model.Pokemon
	mongoDexNum := bson.D{{Key: "Number", Value: dexNum}}
	filter := bson.D{{Key: "$and", Value: bson.A{mongoDexNum}}}

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "Number", Value: 1}})
	res, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Error(err)
	}
	for res.Next(ctx) {
		err := res.Decode(&pokemon)
		if err != nil {
			log.Fatal(err)
		}
		pokemonByDexNum = append(pokemonByDexNum, pokemon)
	}
	return pokemonByDexNum, nil
}

func GetPokedexByOneType(ctx context.Context, collection *mongo.Collection, type1 string) (pokedexByOneType []model.Pokemon, err error) {
	var pokemon model.Pokemon
	MongoType1 := bson.D{{Key: "Element", Value: type1}}
	filter := bson.D{{Key: "$and", Value: bson.A{MongoType1}}}

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "Number", Value: 1}})
	res, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Error(err)
	}
	for res.Next(ctx) {
		err := res.Decode(&pokemon)
		if err != nil {
			log.Fatal(err)
		}
		pokedexByOneType = append(pokedexByOneType, pokemon)
	}
	return pokedexByOneType, nil
}

func GetPokedexByTwoTypes(ctx context.Context, collection *mongo.Collection, type1, type2 string) (pokedexByTwoTypes []model.Pokemon, err error) {
	var pokemon model.Pokemon
	MongoType1 := bson.D{{Key: "Element", Value: type1}}
	MongoType2 := bson.D{{Key: "SecElement", Value: type2}}
	filter := bson.D{{Key: "$and", Value: bson.A{MongoType1, MongoType2}}}

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "Number", Value: 1}})
	res, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Error(err)
	}
	for res.Next(ctx) {
		err := res.Decode(&pokemon)
		if err != nil {
			log.Fatal(err)
		}
		pokedexByTwoTypes = append(pokedexByTwoTypes, pokemon)
	}
	return pokedexByTwoTypes, nil
}
