package services

import (
  "scoreplay/internal/models"
  "scoreplay/internal/repository"
)

func CreateTag(tag *models.Tag) error {
  return repository.CreateTag(tag)
}

func ListTags() ([]models.Tag, error) {
  return repository.ListTags()
}
