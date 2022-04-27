package client

import (
	"context"
	"encoding/json"
	"fmt"
	cfg "github.com/XC-Zero/yinwan/pkg/config"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/olivere/elastic/v7"
	"log"
	"strings"
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

func GetFromIndex(tabler _interface.EsTabler, query elastic.Query, from, size int) (list []interface{}, count int64, err error) {
	highlight := elastic.NewHighlight()

	highlight = highlight.Field("*").PreTags(PRE_TAG).PostTags(POST_TAG)

	log.Println(highlight.Source())
	res, err := ESClient.Search().
		Index(tabler.TableName()).
		From(from).Size(size).
		Highlight(highlight).
		Query(query).
		RestTotalHitsAsInt(true).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, 0, err
	}

	count = res.TotalHits()

	hit := res.Hits.Hits
	for i := range hit {
		var m map[string]interface{}
		hl := hit[i].Highlight
		err = json.Unmarshal(hit[i].Source, &m)
		if err != nil {
			continue
		}
		for s, strList := range hl {
			m[s] = strings.Join(strList, "")
		}
		list = append(list, m)
	}

	return
}

func DeleteIndex(tabler _interface.EsTabler) bool {
	do, err := ESClient.DeleteIndex(tabler.TableName()).Do(context.Background())
	if err != nil {
		return false
	}
	return do.Acknowledged
}
