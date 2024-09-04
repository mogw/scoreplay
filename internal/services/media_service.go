package services

import (
  "errors"
  "scoreplay/internal/elasticsearch"
  "scoreplay/internal/models"
  "scoreplay/internal/repository"
)

func CreateMedia(name string, tags []string, filePath string) (*models.Media, error) {
  media := &models.Media{Name: name, FilePath: filePath}

  // Fetch tags
  var tagModels []models.Tag
  for _, tagID := range tags {
    tag, err := repository.GetTagByID(tagID)
    if err != nil {
      return nil, errors.New("Tag not found")
    }
    tagModels = append(tagModels, *tag)
  }

  media.Tags = tagModels

  if err := repository.CreateMedia(media); err != nil {
    return nil, err
  }

  // Index media in Elasticsearch
  if err := elasticsearch.IndexMedia(media); err != nil {
    return nil, err
  }

  return media, nil
}

func GetMedia(id string) (*models.Media, error) {
  return repository.GetMedia(id)
}
