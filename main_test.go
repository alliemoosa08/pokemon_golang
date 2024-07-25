package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// Test for successful response from PokeAPI for getPokemons
func TestGetPokemonsSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the PokeAPI response
	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon?limit=100",
		httpmock.NewStringResponder(200, `{"results": [{"name": "bulbasaur", "url": "https://pokeapi.co/api/v2/pokemon/1/"}]}`))

	router := mux.NewRouter()
	router.HandleFunc("/pokemons", getPokemons).Methods("GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "bulbasaur")
	assert.Contains(t, w.Body.String(), "https://pokeapi.co/api/v2/pokemon/1/")
}

// Test for network error when fetching data from PokeAPI for getPokemons
func TestGetPokemonsNetworkError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon?limit=100",
		httpmock.NewErrorResponder(errors.New("network error")))

	router := mux.NewRouter()
	router.HandleFunc("/pokemons", getPokemons).Methods("GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "network error")
}

// Test for non-200 response code from PokeAPI for getPokemons
func TestGetPokemonsNon200Response(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon?limit=100",
		httpmock.NewStringResponder(500, "Internal Server Error"))

	router := mux.NewRouter()
	router.HandleFunc("/pokemons", getPokemons).Methods("GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Test for invalid JSON response from PokeAPI for getPokemons
func TestGetPokemonsInvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon?limit=100",
		httpmock.NewStringResponder(200, `invalid json`))

	router := mux.NewRouter()
	router.HandleFunc("/pokemons", getPokemons).Methods("GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid character")
}

// Test for empty results from PokeAPI for getPokemons
func TestGetPokemonsEmptyResults(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon?limit=100",
		httpmock.NewStringResponder(200, `{"results": []}`))

	router := mux.NewRouter()
	router.HandleFunc("/pokemons", getPokemons).Methods("GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]\n", w.Body.String())
}

// Test for successful response from PokeAPI for getPokemonByID
func TestGetPokemonByIDSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon/1",
		httpmock.NewStringResponder(200, `{"name": "bulbasaur", "id": 1}`))

	router := mux.NewRouter()
	router.HandleFunc("/pokemons/{id}", getPokemonByID).Methods("GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "bulbasaur")
	assert.Contains(t, w.Body.String(), "1")
}

// Test for network error when fetching data from PokeAPI for getPokemonByID
func TestGetPokemonByIDNetworkError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon/1",
		httpmock.NewErrorResponder(errors.New("network error")))

	router := mux.NewRouter()
	router.HandleFunc("/pokemons/{id}", getPokemonByID).Methods("GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "network error")
}

// Test for non-200 response code from PokeAPI for getPokemonByID
func TestGetPokemonByIDNon200Response(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon/1",
		httpmock.NewStringResponder(500, "Internal Server Error"))

	router := mux.NewRouter()
	router.HandleFunc("/pokemons/{id}", getPokemonByID).Methods("GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Test for invalid JSON response from PokeAPI for getPokemonByID
func TestGetPokemonByIDInvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon/1",
		httpmock.NewStringResponder(200, `invalid json`))

	router := mux.NewRouter()
	router.HandleFunc("/pokemons/{id}", getPokemonByID).Methods("GET")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid character")
}
