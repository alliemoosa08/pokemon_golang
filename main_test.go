package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemons(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.Default()
	router.GET("/pokemons", gin.WrapF(getPokemons))

	// Create a new HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemons", nil)

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "name")
	assert.Contains(t, w.Body.String(), "url")
	assert.NotEmpty(t, w.Body.String())
}
