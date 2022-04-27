package main

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/es_model"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {
	// 读取配置文件
	config.InitConfiguration()
	// 开协程监听配置文件修改，实现热加载
	go config.ViperMonitor()
	client.InitSystemStorage(config.CONFIG.StorageConfig)
	indexes, err := client.ESClient.CatIndices().Do(context.TODO())
	if err != nil {
		return
	}
	for _, index := range indexes {
		log.Println(index.Index)
	}
	err = client.CreateIndex(es_model.Material{})
	if err != nil {
		panic(err)
	}
	//material := es_model.Material{
	//	RecID:        89764,
	//	MaterialName: "他喵???",
	//	Remark:       "这他喵又不是测试的？？？？？？",
	//}
	//err = client.PutIntoIndex(material)
	//if err != nil {
	//	panic(err)
	//}
	query := elastic.NewMultiMatchQuery("他喵", "rec_id^999", "remark^2", "material_name^10").Operator("or")

	data, count, err := client.GetFromIndex(es_model.Material{}, query, 0, 2)
	if err != nil {
		panic(err)
	}
	log.Println(count)
	log.Printf("%T  \n %+v", data, data)
}
