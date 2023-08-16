package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"net/http"
	"pkdx-api/pkg/db"
	"strconv"
	"time"
)

var (
	MongoDb mongo.Collection
)

// Router
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/pokedex/national", getNationalDex)
	myRouter.HandleFunc("/pokedex/random", getRandomPokemon)
	myRouter.HandleFunc("/pokedex/{dexNum}", getPokemonByDexNum)
	myRouter.HandleFunc("/pokedex/form/{form}", getPokedexByForm)
	myRouter.HandleFunc("/pokedex/type/{type}", getPokedexByOneType)
	myRouter.HandleFunc("/pokedex/type/", getPokedexByTwoTypes)
	myRouter.HandleFunc("/pokedex/region/{region}", getPokedexByRegion)

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

func getPokedexByForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	form := vars["form"]
	pokedexByForm, err := db.GetPokedexByForm(r.Context(), &MongoDb, form)
	if err = json.NewEncoder(w).Encode(pokedexByForm); err != nil {
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

func getPokedexByRegion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	region := vars["region"]
	pokedexByRegion, err := db.GetPokedexByRegion(r.Context(), &MongoDb, region)
	if err = json.NewEncoder(w).Encode(pokedexByRegion); err != nil {
		log.Error(err)
	}
}

func getRandomPokemon(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 989
	randomDexId := rand.Intn(max-min+1) + min
	randomDexIdString := strconv.Itoa(randomDexId)
	switch len(randomDexIdString) {
	case 1:
		randomDexIdString = "000" + randomDexIdString
	case 2:
		randomDexIdString = "00" + randomDexIdString
	case 3:
		randomDexIdString = "0" + randomDexIdString
	}
	randomPokemon, err := db.GetRandomPokemon(r.Context(), &MongoDb, randomDexIdString)
	if err = json.NewEncoder(w).Encode(randomPokemon); err != nil {
		log.Error(err)
	}
}
