package elasticsearch

import (
  "context"
  "scoreplay/internal/models"
  "scoreplay/internal/utils"
)

func IndexMedia(media *models.Media) error {
  _, err := utils.ESClient.Index().
    Index("media").
    Id(media.ID).
    BodyJson(media).
    Do(context.Background())
  return err
}
