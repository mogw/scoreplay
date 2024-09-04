package main

import (
  "log"
  "github.com/golang-migrate/migrate/v4"
  _ "github.com/golang-migrate/migrate/v4/database/postgres"
  _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
  m, err := migrate.New(
    "file://migrations/sql",
    "postgres://scoreplay:password@db:5432/scoreplay?sslmode=disable",
  )
  if err != nil {
    log.Fatalf("Could not initialize migration: %v", err)
  }

  if err := m.Up(); err != nil && err != migrate.ErrNoChange {
    log.Fatalf("Could not run migration: %v", err)
  }

  log.Println("Migration ran successfully")
}
