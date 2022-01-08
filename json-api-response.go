// https://tutorialedge.net/golang/consuming-restful-api-with-go/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// {
	// 	"name":"kanto",
	// 	"region": {
	// 	  "url":"http:\/\/pokeapi.co\/api\/v2\/region\/1\/",
	// 	  "name":"kanto"
	// 	},
	// 	"version_groups":[ ... ]
	// 	],
	// 	"is_main_series":true,
	// 	"descriptions":[ ... ],
	// 	"pokemon_entries":[
	// 	  {
	// 		"entry_number": 1,
	// 		"pokemon_species": {
	// 		  "url":"http:\/\/pokeapi.co\/api\/v2\/pokemon-species\/1\/",
	// 		  "name":"bulbasaur"
	// 		}
	// 	  }
	// 	  ...
	// 	]
	//   }

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		str := fmt.Sprintf("%v - %v", i+1, responseObject.Pokemon[i].Species.Name)
		fmt.Println(str)
	}
}
