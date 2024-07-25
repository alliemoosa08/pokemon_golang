# Pokemon Golang


## Ensure you have the following installed:

Go (1.18+)
`go get` for fetching Go modules

## Installation
Clone the repository:

`git clone https://github.com/alliemoosa08/pokemon_golang.git`
cd <repository-directory>

## Fetch dependencies:

`go mod tidy`

## Running the API

`go run .`
The API will start and listen on 0.0.0.0:8080.

## Access the API:

Get all Pokémon:
`GET http://0.0.0.0:8080/pokemons`

Get Pokémon by ID:
`GET http://0.0.0.0:8080/pokemons/{id}`

## Code Overview

`main.go:` Contains the main application logic including the router setup and API handlers.
`PokemonHandler():` Sets up the routes for Pokémon endpoints.
`cors.AllowAll().Handler(router):` Applies CORS settings to allow all origins.
`http.ListenAndServe:` Starts the HTTP server.