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
	material := es_model.Material{
		RecID:        1,
		MaterialName: "test01x",
		Remark:       "这他喵是测试的！！！",
	}
	err = client.PutIntoIndex(material)
	if err != nil {
		panic(err)
	}
	q := elastic.NewMatchQuery("material_name", "test01x")
	data, count, err := client.GetFromIndex(material, q, 0, 10)
	if err != nil {
		panic(err)
	}
	log.Println(count)
	log.Printf("%T  \n %+v", data, data)
}
