package config

var CONFIG *config

type config struct {
	ApiConfig     ApiConfig     `json:"api_config"`
	ServiceConfig ServiceConfig `json:"service_config"`
	StorageConfig StorageConfig `json:"storage_config"`
	LogConfig     LogConfig     `json:"log_config"`
}

func InitConfig() {

}
