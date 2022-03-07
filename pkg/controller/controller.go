package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"pkdx-api/pkg/model"
)

var (
	MongoDb mongo.Collection
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/pokedex", returnFullPokedex)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func getPokedex(collection *mongo.Collection) (pokedex []model.Pokemon) {
	res, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		logrus.Error(err)
	}
	for res.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var pokemon model.Pokemon
		err := res.Decode(&pokemon)
		if err != nil {
			log.Fatal(err)
		}
		pokedex = append(pokedex, pokemon)
	}
	return
}

func returnFullPokedex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endoint Hit: Full Pokedex")
	pokedex := getPokedex(&MongoDb)
	json.NewEncoder(w).Encode(pokedex)
}
