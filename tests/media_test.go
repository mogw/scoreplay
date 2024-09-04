package tests

import (
  "io"
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"scoreplay/internal/handlers"
)

// Helper function to create a tag and return its ID
func createTag(t *testing.T, router *gin.Engine, tagName string) {
	// Define the tag input
	tag := map[string]string{
		"name": tagName,
	}

	// Convert input to JSON
	jsonData, err := json.Marshal(tag)
	assert.NoError(t, err)

	// Create a new HTTP request to create a tag
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

	// Parse the response to get the tag ID
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	tagID := response["id"]
	return tagID
}

// Test for uploading media with tags and file
func TestUploadMedia(t *testing.T) {
	// Initialize Gin router
	router := gin.Default()

	// Register routes here
	router.POST("/tags", handlers.CreateTag)
	router.POST("/media", handlers.CreateMedia)

	// Create tags first and get their IDs
	tagID1 := createTag(t, router, "TestTag1")
	tagID2 := createTag(t, router, "TestTag2")

	// Create a new multipart request for uploading media
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the "name" field
	err := writer.WriteField("name", "TestMedia")
	assert.NoError(t, err)

	// Add the "tags" field (array of tag IDs)
	tagsField := []int{tagID1, tagID2}
	tagsJson, err := json.Marshal(tagsField)
	assert.NoError(t, err)
	err = writer.WriteField("tags", string(tagsJson))
	assert.NoError(t, err)

	// Add the "file" field (binary file)
	file, err := os.Open("testdata/test_image.jpg") // Ensure you have a test file here
	assert.NoError(t, err)
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	assert.NoError(t, err)

	_, err = io.Copy(part, file)
	assert.NoError(t, err)

	// Close the multipart writer
	err = writer.Close()
	assert.NoError(t, err)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/media", body)
	assert.NoError(t, err)

	// Set the appropriate content type for multipart
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Use httptest to create a response recorder
	rr := httptest.NewRecorder()

	// Simulate the request using Gin's router
	router.ServeHTTP(rr, req)

	// Assert the status code
	assert.Equal(t, http.StatusCreated, rr.Code)
}

// Mock GET request to retrieve all tags
func TestGetMedias(t *testing.T) {
	// Initialize Gin router
	router := gin.Default()

	// Register the route for listing tags
	router.GET("/media", handlers.GetMedia)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/media", nil)
	assert.NoError(t, err)

	// Use httptest to create a response recorder
	rr := httptest.NewRecorder()

	// Simulate the request using Gin's router
	router.ServeHTTP(rr, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, rr.Code)
}
