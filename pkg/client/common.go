package client

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/XC-Zero/yinwan/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"log"
	"strconv"
	"sync"
)

// todo 实际上不应该导出，这样有修改风险，应该应该以单例,后面再改吧
//
//var instance *client
//var once sync.Once
//
//func GetClientInstance()*client{
//    once.Do(func(){
//        instance=client.new()
//    })
//    return instance
//}
var bk *bookNameMap
var once sync.Once

type bookNameMap struct {
	sync.RWMutex
	bookNameMap map[string]BookName
}

func GetBookNameInstance() *bookNameMap {
	once.Do(func() {
		bk.Lock()
		bk = &bookNameMap{bookNameMap: make(map[string]BookName, 0)}
		bk.Unlock()
	})
	return bk
}
func ReadBookMap(key string) (value BookName, ok bool) {
	GetBookNameInstance().RLock()
	value, ok = GetBookNameInstance().bookNameMap[key]
	GetBookNameInstance().RUnlock()
	return
}
func AddBookMap(key string, bn BookName) bool {
	GetBookNameInstance().Lock()
	defer GetBookNameInstance().Unlock()
	if _, ok := GetBookNameInstance().bookNameMap[key]; ok {
		return false
	}
	GetBookNameInstance().bookNameMap[key] = bn
	return true
}
func GetAllBookMap() []BookName {
	var bkList []BookName
	GetBookNameInstance().Lock()
	defer GetBookNameInstance().Unlock()
	bk := GetBookNameInstance().bookNameMap
	for key := range bk {
		bkList = append(bkList, bk[key])
	}
	return bkList

}

var (
	// BookNameMap 账套Map

	RedisClient    *redis.Client
	ESClient       *elastic.Client
	MysqlClient    *gorm.DB
	MinioClient    *minio.Client
	InfluxDBClient *influxdb2.Client
	MongoDBClient  *mongo.Database
	KafkaClient    *sarama.Client
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
	//todo kafka
	//kk, err := InitKafka(config.KafkaConfig)
	//if err != nil {
	//	panic(err)
	//}
	rr, err := InitRedis(config.RedisConfig)
	if err != nil {
		panic(err)
	}
	//todo es
	es, err := InitElasticsearch(config.ESConfig)
	if err != nil {
		panic(err)
	}

	ESClient = es
	RedisClient = rr
	MysqlClient = msy
	MinioClient = mClient
	MongoDBClient = mgClient
	//KafkaClient = kk
}

// Paginate 分页函数 纯 gorm 时在 scope 里调用
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

// MysqlPaginateSql 生成分页sql语句
func MysqlPaginateSql(ctx *gin.Context) string {
	pageNumber := ctx.PostForm("page_number")
	pageSize := ctx.PostForm("page_size")
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

func ESPaginate(ctx *gin.Context, s *elastic.SearchService) *elastic.SearchService {
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
	s.From(int(offset)).Size(int(limit))

	return s
}
