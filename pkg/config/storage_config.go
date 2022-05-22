package config

// StorageConfig 系统存储配置
type StorageConfig struct {
	MysqlConfig   MysqlConfig   `json:"mysql_config" yaml:"mysql_config"`
	MongoDBConfig MongoDBConfig `json:"mongo_db_config" yaml:"mongo_db_config"`
	RedisConfig   RedisConfig   `json:"redis_config" yaml:"redis_config"`
	MinioConfig   MinioConfig   `json:"minio_config" yaml:"minio_config"`
	KafkaConfig   KafkaConfig   `json:"kafka_config" yaml:"kafka_config"`
	InfluxConfig  InfluxConfig  `json:"influx_config" yaml:"influx_config"`
	ESConfig      ESConfig      `json:"es_config" yaml:"es_config"`
}

// BookConfig 账套存储配置
type BookConfig struct {
	BookName      string        `json:"book_name" yaml:"book_name"`
	StorageName   string        `json:"storage_name" yaml:"storage_name"`
	MysqlConfig   MysqlConfig   `json:"mysql_config" yaml:"mysql_config"`
	MinioConfig   MinioConfig   `json:"minio_config" yaml:"minio_config"`
	MongoDBConfig MongoDBConfig `json:"mongo_db_config" yaml:"mongo_db_config"`
}

type MysqlConfig struct {
	Host     string       `json:"host" yaml:"host"`
	Port     string       `json:"port" yaml:"port"`
	UserName string       `json:"user_name" yaml:"user_name"`
	Password string       `json:"password" yaml:"password"`
	DBName   string       `json:"db_name" yaml:"db_name"`
	LogMode  MysqlLogMode `json:"log_mode" yaml:"log_mode"`
}

type ESConfig struct {
	Host                         string `json:"host" yaml:"host"`
	User                         string `json:"user" yaml:"user"`
	Password                     string `json:"password" yaml:"password"`
	ResponseHeaderTimeoutSeconds int    `json:"response_header_timeout_seconds" yaml:"response_header_timeout_seconds"`
}

type InfluxConfig struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Token  string `json:"token" yaml:"token"`
	Bucket string `json:"bucket" yaml:"bucket"`
	Org    string `json:"org" yaml:"org"`
}

type MongoDBConfig struct {
	URL    string `json:"url" yaml:"url"`
	DBName string `json:"db_name" yaml:"db_name"`
}

type MinioConfig struct {
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
	EndPoint  string `json:"end_point" yaml:"end_point"`
	Bucket    string `json:"bucket" yaml:"bucket"`
}

type RedisConfig struct {
	AddrList []string `json:"addr_list" yaml:"addr_list"`
	Password string   `json:"password" yaml:"password"`
}
type KafkaConfig struct {
	AddrList []string `json:"addr_list" yaml:"addr_list"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
}
