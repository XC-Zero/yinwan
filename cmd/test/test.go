package main

import (
	"fmt"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model"
	"log"
	"reflect"
)

func main() {
	//// 读取配置文件
	//config.InitConfiguration()
	//// 开协程监听配置文件修改，实现热加载
	//go config.ViperMonitor()
	//mgClient, err := client.InitMongoDB(config.CONFIG.StorageConfig.MongoDBConfig)
	//if err != nil {
	//	panic(err)
	//}
	//mgClient.Client().Database("hello2")
	//
	//return
	a := reflectNewSlice(model.Staff{})
	log.Println(reflect.TypeOf(a))
}

//反射创建新对象。
func reflectNewSlice(target interface{}) []_interface.ChineseTabler {
	if target == nil {
		fmt.Println("参数不能未空")
	}

	a := reflect.Zero(reflect.ArrayOf(0, reflect.TypeOf(target))).Interface()
	log.Println(reflect.TypeOf(a))
	return reflect.ValueOf(a).Convert(reflect.TypeOf([]_interface.ChineseTabler{})).Interface().([]_interface.ChineseTabler)
}
