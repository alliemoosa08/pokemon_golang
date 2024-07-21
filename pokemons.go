package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokeAPIResponse struct {
	Results []Pokemon `json:"results"`
}

type ErrorHandler struct {
	ErrorDetails string `json:"errorDeatils"`
}

var errorHandler ErrorHandler

func getPokemons(w http.ResponseWriter, _ *http.Request) {
	// Fetch data from PokeAPI
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=100")
	if err != nil {
		errorHandler.ErrorDetails = err.Error()
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorHandler)
		return
	}
	defer resp.Body.Close()

	// Ensure the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(nil)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorHandler.ErrorDetails = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler)
		return
	}

	var pokeAPIResponse PokeAPIResponse
	if err := json.Unmarshal(body, &pokeAPIResponse); err != nil {
		errorHandler.ErrorDetails = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokeAPIResponse.Results)
}

func getPokemonByID(w http.ResponseWriter, r *http.Request) {
	// Extract the Pokemon ID from the request URL
	id := mux.Vars(r)["id"]

	// Fetch data from PokeAPI by ID
	url := "https://pokeapi.co/api/v2/pokemon/" + id

	resp, err := http.Get(url)
	if err != nil {
		errorHandler.ErrorDetails = err.Error()
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorHandler)
		return
	}
	defer resp.Body.Close()

	// Ensure the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(nil)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorHandler.ErrorDetails = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler)
		return
	}

	var pokeDetails map[string]interface{}
	if err := json.Unmarshal(body, &pokeDetails); err != nil {
		errorHandler.ErrorDetails = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorHandler)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokeDetails)
}
