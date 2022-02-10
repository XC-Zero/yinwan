package main

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller"
	"github.com/XC-Zero/yinwan/pkg/client"
)

func main() {
	config.InitConfiguration()
	go config.ViperMonitor()

	client.MysqlClient = client.InitMysqlGormV2(config.CONFIG.StorageConfig.MysqlConfig)

	services_controller.Starter()
	select {}
}
