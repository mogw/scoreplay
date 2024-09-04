package models

type Media struct {
  ID       string `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
  Name     string `json:"name"`
  FilePath string `json:"file_path"`
  Tags     []Tag  `gorm:"many2many:media_tags;" json:"tags"`
}
