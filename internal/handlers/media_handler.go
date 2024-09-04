package handlers

import (
  "net/http"
  "scoreplay/internal/services"
  "scoreplay/internal/utils"

  "github.com/gin-gonic/gin"
)

func CreateMedia(c *gin.Context) {
  name := c.PostForm("name")
  tags := c.PostFormArray("tags")

  file, err := c.FormFile("file")
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
    return
  }

  filePath, err := utils.SaveFile(file)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
    return
  }

  media, err := services.CreateMedia(name, tags, filePath)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusCreated, media)
}

func GetMedia(c *gin.Context) {
  id := c.Param("id")

  media, err := services.GetMedia(id)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, media)
}
