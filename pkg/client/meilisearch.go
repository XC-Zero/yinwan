package client

import (
	"github.com/meilisearch/meilisearch-go"
)

func InitMeilisearch(config meilisearch.ClientConfig) {
	if MeilisearchClient != nil {
		MeilisearchClient = meilisearch.NewClient(config)
	}
}
