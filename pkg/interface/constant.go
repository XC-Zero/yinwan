package _interface

// Constant 自定义类型常量需实现该接口
type Constant interface {
	// Display 获取中文名
	Display() string
}
