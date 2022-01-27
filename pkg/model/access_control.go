package model

import (
	"log"
	"strconv"
	"strings"
)

type ModuleCapability int

const (
	READ ModuleCapability = iota + 620001
	WRITE
	DELETE
	ALL
)

func (m ModuleCapability) Display() string {
	switch m {
	case READ:
		return "查看"
	case WRITE:
		return "修改"
	case DELETE:
		return "删除"
	case ALL:
		return "所有"
	default:
		return "未知"
	}
}

// AccessControl 权限控制 --- 按模块
type AccessControl struct {
	BasicModel
	ModuleID           int    `gorm:"type:int"`
	ModuleName         string `gorm:"type:varchar(50)"`
	ModuleCapabilities string `gorm:"type:varchar(50)"`
}

func (ac AccessControl) SplitCapabilities() (caps []ModuleCapability) {
	cs := ac.ModuleCapabilities
	if cs != "" {
		for _, str := range strings.Split(cs, SPLIT_SYMBOL) {
			val, err := strconv.Atoi(str)
			if err != nil {
				log.Println("转换数据库中 访问模块：功能可选项值 -> 结构体定义常量失败，error is " + err.Error())
				continue
			}
			caps = append(caps, ModuleCapability(val))
		}
	}
	return
}
func (ac *AccessControl) SetModuleCapabilities(caps []ModuleCapability) {
	val := ""
	// 多一个 | 没所谓
	for _, capability := range caps {
		val += strconv.Itoa(int(capability)) + SPLIT_SYMBOL
	}
	ac.ModuleCapabilities = val
}
