package config

import (
	config2 "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/spf13/viper"
)

var CONFIG *config

type config struct {
	ApiConfig      config2.ApiConfig     `json:"api_config"`
	ServiceConfig  config2.ServiceConfig `json:"service_config"`
	StorageConfig  config2.StorageConfig `json:"storage_config"`
	BookNameConfig config2.BookConfig    `json:"book_name_config"`
	LogConfig      config2.LogConfig     `json:"log_config"`
}

func InitConfiguration(configName string, configPaths []string, config interface{}) error {
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.AutomaticEnv()
	for _, configPath := range configPaths {
		vp.AddConfigPath(configPath)
	}

	if err := vp.ReadInConfig(); err != nil {

		panic(err)
	}

	err := vp.Unmarshal(config)
	if err != nil {
		panic(err)
	}

	return nil
}
