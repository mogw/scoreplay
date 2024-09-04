package utils

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
  "scoreplay/internal/models"
)

var DB *gorm.DB

func InitDB() {
  dsn := "host=db user=scoreplay password=password dbname=scoreplay port=5432 sslmode=disable"
  var err error
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalf("Failed to connect to database: %v", err)
  }

  // Enable the uuid-ossp extension
  err = DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
  if err != nil {
    log.Fatalf("Failed to enable uuid-ossp extension: %v", err)
  }

  // Automatically create the schema based on the models
  err = DB.AutoMigrate(&models.Tag{}, &models.Media{})
  if err != nil {
    log.Fatalf("Failed to migrate schema: %v", err)
  }
}
