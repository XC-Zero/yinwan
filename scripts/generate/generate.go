package main

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/client"
)

func main() {
	config.InitConfiguration()
	go config.ViperMonitor()

	client.InitSystemStorage(config.CONFIG.StorageConfig)
	//GenerateMysqlSchema()

}
