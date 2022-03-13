package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"pkdx-api/pkg/model"
)

var (
	MongoDb mongo.Collection
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/pokedex/national", returnFullPokedex)
	myRouter.HandleFunc("/pokedex/{id}", returnSinglePokemon)
	myRouter.HandleFunc("/pokedex/type/{type}", returnAllOfOneType)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func getPokedex(ctx context.Context, collection *mongo.Collection) (pokedex []model.Pokemon) {
	res, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error(err)
	}
	for res.Next(ctx) {
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

//Currently returns the full(national) pokedex
func returnFullPokedex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endoint Hit: Full Pokedex")
	pokedex := getPokedex(r.Context(), &MongoDb)
	if err := json.NewEncoder(w).Encode(pokedex); err != nil {
		log.Error(err)
	}
}

//Options:
//Numbers 001 through 905
func returnSinglePokemon(w http.ResponseWriter, r *http.Request) {
	pokedex := getPokedex(r.Context(), &MongoDb)
	vars := mux.Vars(r)
	key := vars["id"]
	for _, pokemon := range pokedex {
		if pokemon.Id == key {
			if err := json.NewEncoder(w).Encode(pokemon); err != nil {
				log.Error(err)
			}
			fmt.Println("Endpoint Hit: Single Pokemon " + pokemon.Name)
		}
	}
}

//Options:
//Fire, Water, Grass, Electric, Psychic, Flying, Dark, Steel, Fairy,
//Ground, Rock, Normal, Bug, Fighting, Ghost, Ice, Poison, Dragon
func returnAllOfOneType(w http.ResponseWriter, r *http.Request) {
	pokedex := getPokedex(r.Context(), &MongoDb)
	vars := mux.Vars(r)
	key := vars["type"]
	fmt.Println("Endpoint Hit: All Pokemon of One Type " + key)
	for _, pokemon := range pokedex {
		if pokemon.Element == key || pokemon.SecElement == key {
			if err := json.NewEncoder(w).Encode(pokemon); err != nil {
				log.Error(err)
			}

		}
	}
}

//TODO: Create Endpoint for specific pokedexes based on region

//TODO: Create Endpoint for pokemon with specific forms

//TODO: Create Endpoint to return random pokemon
