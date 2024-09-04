package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"scoreplay/internal/handlers"

// Mock POST request to create a tag
func TestCreateTag(t *testing.T) {
	// Initialize Gin router
	router := gin.Default()

	// Register the route for tag creation
	router.POST("/tags", handlers.CreateTag)

	// Define the input
	tag := map[string]string{
		"name": "TestTag",
	}

	// Convert input to JSON
	jsonData, err := json.Marshal(tag)
	assert.NoError(t, err)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/tags", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)

	// Set the header for content type
	req.Header.Set("Content-Type", "application/json")

	// Use httptest to create a response recorder
	rr := httptest.NewRecorder()

	// Simulate the request using Gin's router
	router.ServeHTTP(rr, req)

	// Assert the status code
	assert.Equal(t, http.StatusCreated, rr.Code)
}

// Mock GET request to retrieve all tags
func TestGetTags(t *testing.T) {
	// Initialize Gin router
	router := gin.Default()

	// Register the route for listing tags
	router.GET("/tags", handlers.ListTags)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/tags", nil)
	assert.NoError(t, err)

	// Use httptest to create a response recorder
	rr := httptest.NewRecorder()

	// Simulate the request using Gin's router
	router.ServeHTTP(rr, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, rr.Code)
}
