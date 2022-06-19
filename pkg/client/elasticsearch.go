package client

import (
	"context"
	"encoding/json"
	"fmt"
	cfg "github.com/XC-Zero/yinwan/pkg/config"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"log"
	"reflect"
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

// PutIntoIndex 添加数据
func PutIntoIndex(tabler _interface.EsTabler) error {
	_, err := ESClient.Index().
		Index(tabler.TableName()).
		BodyJson(tabler.ToESDoc()).
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

// GetFromIndex 获取数据， 所有字段均带高亮（指定字段失败）
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
	l := res.Each(reflect.TypeOf(tabler))
	log.Printf("%+v", l)
	hit := res.Hits.Hits
	for i := range hit {
		var m map[string]interface{}
		hl := hit[i].Highlight
		err = json.Unmarshal(hit[i].Source, &m)
		if err != nil {
			logger.Error(errors.WithStack(err), "Unmarshal Error! ")
			continue
		}
		for s, strList := range hl {
			m[s] = strings.Join(strList, "")
		}
		list = append(list, m)
	}

	return
}

// DeleteFromIndex 删除数据
func DeleteFromIndex(tabler _interface.EsTabler, recID *int, ctx context.Context) error {
	bookName := ctx.Value("book_name").(string)
	b, ok := ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}

	if recID == nil || b.StorageName == "" {
		return errors.New("缺少主键！")
	}
	do, err := ESClient.DeleteByQuery(tabler.TableName()).Query(elastic.NewTermQuery("rec_id", recID)).Do(context.Background())
	if err != nil {
		return err
	}
	log.Println(do.Total)
	return nil
}

// DeleteIndex 删除索引！！！ 谨慎使用
func DeleteIndex(tabler _interface.EsTabler) bool {
	do, err := ESClient.DeleteIndex(tabler.TableName()).Do(context.Background())
	if err != nil {
		return false
	}
	return do.Acknowledged
}

// UpdateIntoIndex 更新内容
func UpdateIntoIndex(tabler _interface.EsTabler, recID *int, ctx context.Context, script *elastic.Script) error {
	bookName := ctx.Value("book_name").(string)
	b, ok := ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}

	if recID == nil || b.StorageName == "" {
		return errors.New("缺少主键！")
	}

	if recID == nil {
		return errors.New("缺少主键！")
	}
	do, err := ESClient.UpdateByQuery(tabler.TableName()).Query(elastic.NewTermQuery("rec_id", recID)).Script(script).Refresh("true").Do(context.Background())
	if err != nil {
		return err
	}
	log.Println(do.Total)
	return nil
}
