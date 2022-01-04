package config

// StorageConfig 存储配置
type StorageConfig struct {
	MysqlConfig   MysqlConfig   `json:"mysql_config"`
	MongoDBConfig MongoDBConfig `json:"mongo_db_config"`
	RedisConfig   RedisConfig   `json:"redis_config"`
	MinioConfig   MinioConfig   `json:"minio_config"`
	KafkaConfig   KafkaConfig   `json:"kafka_config"`
	InfluxConfig  InfluxConfig  `json:"influx_config"`
}

type MysqlConfig struct {
	Host     string       `json:"host"`
	Port     string       `json:"port"`
	UserName string       `json:"user_name"`
	Password string       `json:"password"`
	DBName   string       `json:"db_name"`
	LogMode  MysqlLogMode `json:"log_mode"`
}

type ESConfig struct {
	Host                         string `json:"host"`
	User                         string `json:"user"`
	Password                     string `json:"password"`
	ResponseHeaderTimeoutSeconds int    `json:"response_header_timeout_seconds"`
}

type InfluxConfig struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Token  string `json:"token"`
	Bucket string `json:"bucket"`
	Org    string `json:"org"`
}

type MongoDBConfig struct {
	Host   string `json:"host"`
	DBName string `json:"db_name"`
}

type MinioConfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	EndPoint  string `json:"end_point"`
	Bucket    string `json:"bucket"`
}

type RedisConfig struct {
	AddrList []string `json:"addr_list"`
	Password string   `json:"password"`
}
type KafkaConfig struct {
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
}
