package config

import config2 "github.com/XC-Zero/yinwan/pkg/config"

var CONFIG *config

type config struct {
	ApiConfig     config2.ApiConfig     `json:"api_config"`
	ServiceConfig config2.ServiceConfig `json:"service_config"`
	StorageConfig config2.StorageConfig `json:"storage_config"`
	LogConfig     config2.LogConfig     `json:"log_config"`
}

func InitConfig() {

}
