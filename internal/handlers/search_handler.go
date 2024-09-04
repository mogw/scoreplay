package handlers

import (
  "net/http"
  "scoreplay/internal/services"

  "github.com/gin-gonic/gin"
)

func SearchMedia(c *gin.Context) {
  query := c.Query("q")
  if query == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
    return
  }

  results, err := services.SearchMedia(query)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, results)
}
