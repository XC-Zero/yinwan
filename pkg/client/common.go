package client

import (
	"github.com/Shopify/sarama"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v7"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/meilisearch/meilisearch-go"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient
	ESClient           *elasticsearch.Client
	MeilisearchClient  *meilisearch.Client
	MysqlClient        *gorm.DB
	InfluxDBClient     *influxdb2.Client
	MongoDBClient      *mongo.Database
	KafkaClient        *sarama.Client
)
