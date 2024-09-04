package repository

import (
  "log"

  "scoreplay/internal/models"
  "scoreplay/internal/utils"
)

func CreateTag(tag *models.Tag) error {
  return utils.DB.Create(tag).Error
}

func ListTags() ([]models.Tag, error) {
  var tags []models.Tag
  if err := utils.DB.Find(&tags).Error; err != nil {
    return nil, err
  }
  return tags, nil
}

func GetTagByID(id string) (*models.Tag, error) {
  var tag models.Tag
  err := utils.DB.First(&tag, "id = ?", id).Error
  if err != nil {
    log.Printf("Error fetching tag with ID %s: %v", id, err)
    return nil, err
  }

  return &tag, nil
}
