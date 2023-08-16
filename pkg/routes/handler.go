package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"pkdx-api/pkg/db"
)

var (
	MongoDb mongo.Collection
)

// Router
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/pokedex/national", getNationalDex)
	myRouter.HandleFunc("/pokedex/{dexNum}", getPokemonByDexNum)
	myRouter.HandleFunc("/pokedex/{type}", getPokedexByOneType)
	myRouter.HandleFunc("/pokedex", getPokedexByTwoTypes)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func getNationalDex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endoint Hit: Full Pokedex")
	nationalDex, err := db.GetPokedex(r.Context(), &MongoDb)
	if err = json.NewEncoder(w).Encode(nationalDex); err != nil {
		log.Error(err)
	}
}

func getPokemonByDexNum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dexNum := vars["dexNum"]
	pokemonByDexNum, err := db.GetPokemonByDexNum(r.Context(), &MongoDb, dexNum)
	if err = json.NewEncoder(w).Encode(pokemonByDexNum); err != nil {
		log.Error(err)
	}
}

func getPokedexByOneType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	type1 := vars["type"]
	pokedexByOneType, err := db.GetPokedexByOneType(r.Context(), &MongoDb, type1)
	if err = json.NewEncoder(w).Encode(pokedexByOneType); err != nil {
		log.Error(err)
	}
}

func getPokedexByTwoTypes(w http.ResponseWriter, r *http.Request) {
	type1, ok := r.URL.Query()["type1"]
	if !ok {
		errors.New(`type1 and type2 are required parameters`)
	}
	type2, ok := r.URL.Query()["type2"]
	if !ok {
		errors.New(`type1 and type2 are required parameters`)
	}
	pokedexByTwoTypes, err := db.GetPokedexByTwoTypes(r.Context(), &MongoDb, type1[0], type2[0])
	if err = json.NewEncoder(w).Encode(pokedexByTwoTypes); err != nil {
		log.Error(err)
	}
}

//TODO: Create Endpoint for specific pokedexes based on region

//TODO: Create Endpoint for pokemon with specific forms

//TODO: Create Endpoint to return random pokemon
