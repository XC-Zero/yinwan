package client

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/XC-Zero/yinwan/pkg/config"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"log"
	"strconv"
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
	log.Println("????")
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
	//kk, err := InitKafka(config.KafkaConfig)
	//if err != nil {
	//	panic(err)
	//}
	rr, err := InitRedis(config.RedisConfig)
	if err != nil {
		panic(err)
	}
	//es, err := InitElasticsearch(config.ESConfig)
	if err != nil {
		panic(err)
	}

	//ESClient = es
	RedisClient = rr
	MysqlClient = msy
	MinioClient = mClient
	MongoDBClient = mgClient
	//KafkaClient = kk
}

// Paginate 分页函数
func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageNumber := ctx.PostForm("page_number")
		pageSize := ctx.PostForm("page_size")
		log.Println(pageNumber, pageSize)
		n, limit := 0, 0
		if pn, err := strconv.Atoi(pageNumber); err != nil {
			n = 1
		} else {
			n = pn
		}
		if ps, err := strconv.Atoi(pageSize); err != nil {
			limit = 5
		} else {
			limit = ps
		}

		offset := (n - 1) * limit
		log.Println(offset, limit)
		return db.Offset(offset).Limit(limit)
	}
}

// MysqlPaginateSql 生成分页sql
func MysqlPaginateSql(ctx *gin.Context) string {
	pageNumber := ctx.PostForm("page_number")
	pageSize := ctx.PostForm("page_size")
	log.Println(pageNumber, pageSize)
	n, limit := 0, 0
	if pn, err := strconv.Atoi(pageNumber); err != nil {
		n = 1
	} else {
		n = pn
	}
	if ps, err := strconv.Atoi(pageSize); err != nil {
		limit = 10
	} else {
		limit = ps
	}

	offset := (n - 1) * limit
	return fmt.Sprintf(" limit %d,%d ", offset, limit)
}

func MongoPaginate(ctx *gin.Context, options *options.FindOptions) *options.FindOptions {
	pageNumber := ctx.PostForm("page_number")
	pageSize := ctx.PostForm("page_size")
	log.Println(pageNumber, pageSize)
	var n, limit int64 = 0, 0
	if pn, err := strconv.Atoi(pageNumber); err != nil {
		n = 1
	} else {
		n = int64(pn)
	}
	if ps, err := strconv.Atoi(pageSize); err != nil {
		limit = 10
	} else {
		limit = int64(ps)
	}

	offset := (n - 1) * limit

	options.Limit = &limit
	options.Skip = &offset
	return options
}
