package main

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/client"
	"testing"
)

func Test(t *testing.T) {
	// 读取配置文件
	config.InitConfiguration()
	// 开协程监听配置文件修改，实现热加载
	go config.ViperMonitor()
	client.InitSystemStorage(config.CONFIG.StorageConfig)
}

// TODO
//  1.CRUD方法上都加上 Context 传递当前账套
//  2.common.HarvestClientFromGinContext()  添加账套,都在请求头里读好了
