package main

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/client"
	"log"
)

func main() {
	config.InitConfiguration()
	go config.ViperMonitor()

	//client.MysqlClient = client.InitMysqlGormV2(config.CONFIG.StorageConfig.MysqlConfig)
	minio, err := client.InitMinio(config.CONFIG.StorageConfig.MinioConfig)
	if err != nil {
		panic(err)
	}
	exists, err := minio.BucketExists(context.TODO(), "test01x")
	if err != nil {
		panic(err)
	}
	log.Println(exists)

	//services_controller.Starter()
	select {}
}
