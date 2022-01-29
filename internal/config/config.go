package config

import (
	"flag"
	"fmt"
	config2 "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

var configName = flag.String("c", "config", "读入config文件")
var configPath = flag.String("p", "./configs", "config所在目录")
var CONFIG *config
var GlobalViper *viper.Viper

func init() {
	CONFIG = new(config)
	GlobalViper = viper.New()
}

type config struct {
	ApiConfig      config2.ApiConfig     `json:"api_config"`
	ServiceConfig  config2.ServiceConfig `json:"service_config"`
	StorageConfig  config2.StorageConfig `json:"storage_config"`
	BookNameConfig []config2.BookConfig  `json:"book_name_config"`
	LogConfig      config2.LogConfig     `json:"log_config"`
}

func setTagName(d *mapstructure.DecoderConfig) {
	d.TagName = "json"
}

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

func ViperMonitor() {
	flag.Parse()
	GlobalViper.WatchConfig()
	GlobalViper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info(fmt.Sprintf("检测到配置文件 %s, 开始重新读取配置文件！", e.String()))
		InitConfiguration()
	})
}
