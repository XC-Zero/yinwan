package main

import (
	"fmt"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"log"
	"reflect"
)

func main() {

}

type Person interface {
	Walk()
}

type Man struct {
	Name string
	Age  int32
}

func (m Man) Walk() {
	fmt.Println("man walk")
}

// 检查Person接口对应iface结构体的data成员
//func PrintData(p interface{}) {
//	pd := (**[]Man)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(p)))
//	var m []Man
//	fmt.Printf("iface &data:%p, data:%p, *data:%+v\n", pd, *pd, m)
//	m = **pd
//
//}

//反射创建新对象。
func reflectNewSlice(target interface{}) []_interface.ChineseTabler {
	if target == nil {
		fmt.Println("参数不能未空")
	}

	a := reflect.Zero(reflect.ArrayOf(0, reflect.TypeOf(target))).Interface()
	log.Println(reflect.TypeOf(a))
	return reflect.ValueOf(a).Convert(reflect.TypeOf([]_interface.ChineseTabler{})).Interface().([]_interface.ChineseTabler)
}
