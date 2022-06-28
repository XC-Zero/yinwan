package main

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller"
	"github.com/XC-Zero/yinwan/pkg/client"
)

// @title HSH
// @version 1.0
func main() {
	// 读取配置文件
	config.InitConfiguration()
	// 开协程监听配置文件修改，实现热加载
	go config.ViperMonitor()
	client.InitSystemStorage(config.CONFIG.StorageConfig)
	services_controller.Starter()

}
