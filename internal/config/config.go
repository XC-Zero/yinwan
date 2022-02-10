package config

import (
	"flag"
	"fmt"
	config2 "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

var CONFIG *config
var GlobalViper *viper.Viper

var configName = flag.String("c", "config", "读入config文件")
var configPath = flag.String("p", "./configs", "config所在目录")

func init() {
	CONFIG = new(config)
	// 使用全局 viper 的实例化对象，保证监听和初始化时使用同一个 viper 对象
	GlobalViper = viper.New()
}

type config struct {
	ApiConfig      config2.ApiConfig     `json:"api_config"`
	ServiceConfig  config2.ServiceConfig `json:"service_config"`
	StorageConfig  config2.StorageConfig `json:"storage_config"`
	BookNameConfig []config2.BookConfig  `json:"book_name_config"`
	LogConfig      config2.LogConfig     `json:"log_config"`
}

// 设置 config 对应的结构体的 tag
func setTagName(d *mapstructure.DecoderConfig) {
	d.TagName = "json"
}

// InitConfiguration 初始化配置
// 使用 viper 读取指定路径的指定文件，将内容映射到结构体里
func InitConfiguration() {
	GlobalViper.SetConfigName(*configName)
	GlobalViper.AutomaticEnv()
	GlobalViper.AddConfigPath(*configPath)

	if err := GlobalViper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := GlobalViper.Unmarshal(CONFIG, setTagName)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", CONFIG)
	return
}

// ViperMonitor 使用 viper 监听指定文件
func ViperMonitor() {
	flag.Parse()
	GlobalViper.WatchConfig()
	GlobalViper.OnConfigChange(func(e fsnotify.Event) {
		log.Println(fmt.Sprintf("检测到配置文件 %s, 开始重新读取配置文件！", e.String()))
		//logger.Info(fmt.Sprintf("检测到配置文件 %s, 开始重新读取配置文件！", e.String()))
		InitConfiguration()
	})
}

func SaveConfig(key string, value interface{}) error {

	GlobalViper.Set(key, value)
	err := GlobalViper.WriteConfig()

	return err

}
