package main

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/internal/services/system"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/minio/minio-go/v7"
	"log"
	"os"
)

func main() {
	config.InitConfiguration()
	go config.ViperMonitor()
	client.InitSystemStorage(config.CONFIG.StorageConfig)
	if system.AddBookName("测试一号！") {
		bookname := client.BookNameMap["测试一号！"]
		log.Println(bookname.BookName)
		file, err := os.Open("./configs/config.yml")
		if err != nil {
			panic(err)
		}
		info, err := bookname.MinioClient.PutObject(context.TODO(), bookname.StorageName, "config.yml", file, -1, minio.PutObjectOptions{})
		if err != nil {
			panic(err)
		}
		log.Println(info)
	}
}
