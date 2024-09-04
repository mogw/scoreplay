package services

import (
  "context"
  "encoding/json"
  "scoreplay/internal/models"
  "scoreplay/internal/utils"

  "github.com/olivere/elastic/v7"
)

func SearchMedia(query string) ([]models.Media, error) {
  var results []models.Media

  searchResult, err := utils.ESClient.Search().
    Index("media").
    Query(elastic.NewMultiMatchQuery(query, "name", "tags.name")).
    Do(context.Background())
  if err != nil {
    return nil, err
  }

  for _, hit := range searchResult.Hits.Hits {
    var media models.Media
    err := json.Unmarshal(hit.Source, &media)
    if err != nil {
      return nil, err
    }
    results = append(results, media)
  }

  return results, nil
}
