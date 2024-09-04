package handlers

import (
  "net/http"
  "scoreplay/internal/models"
  "scoreplay/internal/services"

  "github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) {
  var tag models.Tag
  if err := c.ShouldBindJSON(&tag); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if err := services.CreateTag(&tag); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusCreated, tag)
}

func ListTags(c *gin.Context) {
  tags, err := services.ListTags()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, tags)
}
