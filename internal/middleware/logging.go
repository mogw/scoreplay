package middleware

import (
  "github.com/gin-gonic/gin"
  "log"
  "time"
)

func LoggingMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    startTime := time.Now()

    c.Next()

    latency := time.Since(startTime)
    log.Printf("Request: %s %s | Latency: %s", c.Request.Method, c.Request.URL, latency)
  }
}
