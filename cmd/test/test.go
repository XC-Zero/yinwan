package main

import (
	"fmt"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/system"
	"log"
	"math/rand"
	"time"
)

func main() {
	config.InitConfiguration()
	// 开协程监听配置文件修改，实现热加载
	go config.ViperMonitor()
	client.InitSystemStorage(config.CONFIG.StorageConfig)
	services_controller.Starter()

	select {}

}
