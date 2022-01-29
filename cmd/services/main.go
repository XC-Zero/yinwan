package main

import "github.com/XC-Zero/yinwan/internal/config"

func main() {
	// 读取配置文件
	config.InitConfiguration()
	// 开协程监听配置文件修改，实现热加载
	go config.ViperMonitor()
	select {}
}
