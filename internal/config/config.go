package config

import (
	"flag"
	"fmt"
	config2 "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/ghodss/yaml"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"reflect"
)

var CONFIG *config
var globalViper *viper.Viper

var configName = flag.String("c", "config", "读入config文件")
var configPath = flag.String("p", "./configs", "config所在目录")

func init() {
	CONFIG = new(config)
	// 使用全局 viper 的实例化对象，保证监听和初始化时使用同一个 viper 对象
	globalViper = viper.New()
}

type config struct {
	ApiConfig      config2.ApiConfig     `json:"api_config" yaml:"api_config"`
	ServiceConfig  config2.ServiceConfig `json:"service_config" yaml:"service_config"`
	StorageConfig  config2.StorageConfig `json:"storage_config" yaml:"storage_config"`
	BookNameConfig []config2.BookConfig  `json:"book_name_config" yaml:"book_name_config"`
	LogConfig      config2.LogConfig     `json:"log_config" yaml:"log_config"`
}

// 设置 config 对应的结构体的 tag
func setTagName(d *mapstructure.DecoderConfig) {
	d.TagName = "json"
}

// InitConfiguration 初始化配置
// 使用 viper 读取指定路径的指定文件，将内容映射到结构体里
func InitConfiguration() {
	globalViper.SetConfigName(*configName)
	globalViper.AutomaticEnv()
	globalViper.AddConfigPath(*configPath)

	if err := globalViper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := globalViper.Unmarshal(CONFIG, setTagName)
	if err != nil {
		panic(err)
	}
	return
}

// ViperMonitor 使用 viper 监听指定文件
func ViperMonitor() {
	flag.Parse()
	globalViper.WatchConfig()
	globalViper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info(fmt.Sprintf("检测到配置文件 \n（路径：%s）\n 开始重新读取配置文件！", e.String()))
		InitConfiguration()
	})
}

// SaveConfig 因为 viper.set 不会自动 set json tag 中内容，
// 所以得自己转换一遍 set 进去
// issue: 无法将数组 set 进去
//
// Deprecated: 已弃用
func SaveConfig(key string, value interface{}) error {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		list, _ := value.([]config2.BookConfig)
		var resultList []map[string]interface{}
		for i := range list {
			mm := make(map[string]interface{}, 0)
			err := UnfoldMapToProperties(list[i], key, mm)
			if err != nil {
				return err
			}
			resultList = append(resultList, mm)
		}
		for i := range resultList {
			for key, value := range resultList[i] {
				globalViper.Set(key, value)
			}
		}
	} else {
		log.Println("+++++++++++++++++}")
		mm := make(map[string]interface{}, 0)
		err := UnfoldMapToProperties(value, key, mm)
		if err != nil {
			return err
		}
		for s, i := range mm {
			globalViper.Set(s, i)
		}
	}
	err := globalViper.WriteConfig()
	return err
}

func SaveBookConfig(value interface{}) error {
	marshal, err := yaml.Marshal(value)
	if err != nil {
		return err
	}
	err = writeToFile("./configs/config.yml", marshal)

	if err != nil {
		return err
	}
	return nil
}

func writeToFile(fileName string, content []byte) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	} else {
		//// offset
		//err := os.Truncate(fileName, 0)
		//if err != nil {
		//	return err
		//} //clear
		n, _ := f.Seek(0, io.SeekEnd)
		_, err = f.WriteAt(content, n)
		defer f.Close()
	}
	return nil
}

// UnfoldMapToProperties 展开嵌套结构体为 Properties 结构
// 必须是结构体，否则报错！
//
func UnfoldMapToProperties(val interface{}, prefix string, result map[string]interface{}) error {
	objVal, objType := reflect.ValueOf(val), reflect.TypeOf(val)
	if objType.Kind() != reflect.Struct {
		return errors.New("Type of value  is not struct! ")
	}
	for i := 0; i < objVal.NumField(); i++ {
		name := objType.Field(i).Tag.Get("json")
		kind := objType.Field(i).Type.Kind()
		totalName := prefix + "." + name
		if kind == reflect.Map || kind == reflect.Struct {
			err := UnfoldMapToProperties(objVal.Field(i).Interface(), totalName, result)
			if err != nil {
				return err
			}
		} else {
			result[totalName] = objVal.Field(i).Interface()
		}
	}
	return nil
}
