package client

import (
	"github.com/Shopify/sarama"
	"github.com/XC-Zero/yinwan/pkg/config"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v7"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	// BookNameMap 账套Map
	BookNameMap = make(map[string]BookName, 0)

	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient
	ESClient           *elasticsearch.Client
	MysqlClient        *gorm.DB
	MinioClient        *minio.Client
	InfluxDBClient     *influxdb2.Client
	MongoDBClient      *mongo.Database
	KafkaClient        *sarama.Client
)

// InitSystemStorage 初始化系统配置
func InitSystemStorage(config config.StorageConfig) {
	msy, err := InitMysqlGormV2(config.MysqlConfig)
	if err != nil {
		panic(err)
	}
	mClient, err := InitMinio(config.MinioConfig)
	if err != nil {
		panic(err)
	}
	mgClient, err := InitMongoDB(config.MongoDBConfig)
	if err != nil {
		panic(err)
	}
	kk, err := InitKafka(config.KafkaConfig)
	if err != nil {
		panic(err)
	}

	MysqlClient = msy
	MinioClient = mClient
	MongoDBClient = mgClient
	KafkaClient = kk
}
