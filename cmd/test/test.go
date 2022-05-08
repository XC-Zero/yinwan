package main

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/minio/minio-go/v7"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	config.InitConfiguration()
	// 开协程监听配置文件修改，实现热加载
	go config.ViperMonitor()
	client.InitSystemStorage(config.CONFIG.StorageConfig)

	file, err := os.Open("D:\\Administrator\\Documents\\6.jpg")
	if err != nil {
		panic(err)
	}
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	objectName := "PIC_" + strconv.FormatInt(time.Now().Unix(), 10)

	info, err := client.MinioClient.PutObject(
		context.Background(),
		"images", objectName, file, stat.Size(), minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", info)
}
