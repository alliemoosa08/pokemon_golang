package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var router = mux.NewRouter().StrictSlash(true)

func main() {
	PokemonHandler()
	var portNumber = "8080"
	fmt.Println("Listening on IP 0.0.0.0:" + portNumber)

	handle := cors.AllowAll().Handler(router)
	err := http.ListenAndServe("0.0.0.0:"+portNumber, handle)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}

func PokemonHandler() {
	router.HandleFunc("/pokemons", getPokemons).Methods("GET")
	router.HandleFunc("/pokemons/{id}", getPokemonByID).Methods("GET")
}
