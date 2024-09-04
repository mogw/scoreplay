package repository

import (
    "scoreplay/internal/models"
    "scoreplay/internal/utils"
)

func CreateMedia(media *models.Media) error {
  return utils.DB.Create(media).Error
}

func GetMedia(id string) (*models.Media, error) {
  var media models.Media
  if err := utils.DB.Preload("Tags").First(&media, "id = ?", id).Error; err != nil {
    return nil, err
  }
  return &media, nil
}
