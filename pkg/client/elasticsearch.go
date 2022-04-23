package client

import (
	"context"
	"fmt"
	cfg "github.com/XC-Zero/yinwan/pkg/config"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/olivere/elastic/v7"
)

// InitElasticsearch ...
func InitElasticsearch(config cfg.ESConfig) (*elastic.Client, error) {
	esClient, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.Host),
		elastic.SetBasicAuth(config.User, config.Password),
	)
	if err != nil {
		panic(err)
	}

	info, code, err := esClient.Ping(config.Host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esVersion, err := esClient.ElasticsearchVersion(config.Host)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esVersion)

	return esClient, nil
}

func CreateIndex(indexName string, mapping ElasticSearchMapping) error {
	exists, err := ESClient.IndexExists(indexName).Do(context.Background())
	if err != nil {
		return err
	}
	if !exists {
		_, err := ESClient.CreateIndex(indexName).
			BodyJson(mapping).
			IncludeTypeName(true).
			Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func PutIntoIndex(tabler _interface.ChineseTabler) error {
	_, err := ESClient.Index().
		Index(tabler.TableName()).
		BodyJson(tabler).
		Do(context.Background())
	if err != nil {
		return err
	}
	_, err = ESClient.Flush().Index(tabler.TableName()).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func GetDataFromIndex(tabler _interface.ChineseTabler, queryString string) *elastic.SearchResult {
	query := elastic.NewTermQuery()
	do, err := ESClient.Search().Index(tabler.TableName()).Query().
		Sort().
		From().
		Size().
		Pretty(true).Do()
	if err != nil {
		return nil
	}
}
