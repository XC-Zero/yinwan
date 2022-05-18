package main

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
)

func main() {
	config.InitConfiguration()
	go config.ViperMonitor()

	client.InitSystemStorage(config.CONFIG.StorageConfig)
	InitMysqlList()
	//err := GenerateSystemMysqlTables(client.MysqlClient)
	//if err != nil {
	//	panic(err)
	//}
	//DropESIndex()
	//GenerateESIndex()
	client.MysqlClient.AutoMigrate(&mysql_model.Provider{})

}
