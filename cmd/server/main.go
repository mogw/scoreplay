package main

import (
  "github.com/gin-gonic/gin"
  "scoreplay/internal/handlers"
  "scoreplay/internal/middleware"
  "scoreplay/internal/utils"
)

func main() {
  utils.InitDB()
  utils.InitElasticsearch()

  r := gin.Default()

  // Apply middleware
  r.Use(middleware.LoggingMiddleware())

  // Routes
  r.POST("/tags", handlers.CreateTag)
  r.GET("/tags", handlers.ListTags)
  r.POST("/media", handlers.CreateMedia)
  r.GET("/media/:id", handlers.GetMedia)

  r.Run(":8080") // Listen on port 8080
}
