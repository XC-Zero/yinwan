package main

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/client"
)

func main() {
	// 读取配置文件
	config.InitConfiguration()
	// 开协程监听配置文件修改，实现热加载
	go config.ViperMonitor()
	mgClient, err := client.InitMongoDB(config.CONFIG.StorageConfig.MongoDBConfig)
	if err != nil {
		panic(err)
	}
	mgClient.Client().Database("hello2")

	return
}
