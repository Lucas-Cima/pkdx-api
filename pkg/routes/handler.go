package routes

import (
	"encoding/json"
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
	myRouter.HandleFunc("/pokedex/national", returnFullPokedex)
	myRouter.HandleFunc("/pokedex/{dexnum}", returnSinglePokemon)
	myRouter.HandleFunc("/pokedex/type/{type}", returnAllOfOneType)
	myRouter.HandleFunc("/pokedex/type1={type1}&type2={type2}", returnTypeSearch)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

//QuerySinglePokemon

// Currently returns the full(national) pokedex
func returnFullPokedex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endoint Hit: Full Pokedex")
	pokedex := db.GetPokedex(r.Context(), &MongoDb)
	if err := json.NewEncoder(w).Encode(pokedex); err != nil {
		log.Error(err)
	}
}

// Options:
// Numbers 001 through 905
func returnSinglePokemon(w http.ResponseWriter, r *http.Request) {
	pokedex := db.GetPokedex(r.Context(), &MongoDb)
	vars := mux.Vars(r)
	key := vars["dexnum"]
	for _, pokemon := range pokedex {
		if pokemon.DexNum == key {
			if err := json.NewEncoder(w).Encode(pokemon); err != nil {
				log.Error(err)
			}
			fmt.Println("Endpoint Hit: Single Pokemon " + pokemon.Name)
		}
	}
}

// Options:
// Fire, Water, Grass, Electric, Psychic, Flying, Dark, Steel, Fairy,
// Ground, Rock, Normal, Bug, Fighting, Ghost, Ice, Poison, Dragon
func returnAllOfOneType(w http.ResponseWriter, r *http.Request) {
	pokedex := db.GetPokedex(r.Context(), &MongoDb)
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

func returnTypeSearch(w http.ResponseWriter, r *http.Request) {
	pokedex := db.GetPokedex(r.Context(), &MongoDb)
	vars := mux.Vars(r)
	type1 := vars["type1"]
	type2 := vars["type2"]
	for _, pokemon := range pokedex {
		if pokemon.Element == type1 || pokemon.SecElement == type2 {
			if err := json.NewEncoder(w).Encode(pokemon); err != nil {
				log.Error(err)
			}
		}
	}

}

//TODO: Create Endpoint for specific pokedexes based on region

//TODO: Create Endpoint for pokemon with specific forms

//TODO: Create Endpoint to return random pokemon
