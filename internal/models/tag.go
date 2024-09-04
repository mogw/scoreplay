package models

type Tag struct {
  ID   string `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
  Name string `gorm:"uniqueIndex" json:"name"`
}
