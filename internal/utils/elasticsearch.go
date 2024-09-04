package utils

import (
  "github.com/olivere/elastic/v7"
  "log"
)

var ESClient *elastic.Client

func InitElasticsearch() {
  var err error
  ESClient, err = elastic.NewClient(
    elastic.SetURL("http://elasticsearch:9200"),
    elastic.SetSniff(false),
  )
  if err != nil {
    log.Fatalf("Error initializing Elasticsearch: %v", err)
  }
}
