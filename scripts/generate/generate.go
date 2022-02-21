package main

import (
	"github.com/XC-Zero/yinwan/internal/config"
)

func main() {
	config.InitConfiguration()
	go config.ViperMonitor()
	// todo 初始化配置文件
	//client.MysqlClient = client.InitMysqlGormV2(config.CONFIG.StorageConfig.MysqlConfig)
	//GenerateMysqlSchema()

}
