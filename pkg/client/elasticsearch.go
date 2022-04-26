package client

import (
	"context"
	"fmt"
	cfg "github.com/XC-Zero/yinwan/pkg/config"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
)

const (
	PRE_TAG  = "<hl>"
	POST_TAG = "</hl>"
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

func CreateIndex(model _interface.EsTabler) error {
	exists, err := ESClient.IndexExists(model.TableName()).Do(context.Background())
	if err != nil {
		return err
	}
	if !exists {
		_, err := ESClient.CreateIndex(model.TableName()).
			BodyJson(model.Mapping()).
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

func GetFromIndex(tabler _interface.EsTabler, query elastic.Query, from, size int, highlightFields ...string) (list []interface{}, count int64, err error) {
	highlight := elastic.NewHighlight()

	for _, s := range highlightFields {
		highlight.Field(s)
	}
	highlight = highlight.PreTags(PRE_TAG).PostTags(POST_TAG)

	log.Println(highlight.Source())
	res, err := ESClient.Search().
		Index(tabler.TableName()).
		From(from).Size(size).
		Query(query).
		Highlight(highlight).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, 0, err
	}

	count = res.TotalHits()

	// TODO WTF? Why can not highlight keyword ?????
	for i := range res.Hits.Hits {
		log.Println(*res.Hits.Hits[i])

	}
	list = res.Each(reflect.TypeOf(tabler))

	return
}
