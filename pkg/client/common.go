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

type bookNameMap struct {
	sync.RWMutex
	bookNameMap map[string]BookName
}

func initBookName() {
	bk = &bookNameMap{bookNameMap: make(map[string]BookName, 0)}
	bk.bookNameMap["basic"] = BookName{
		BookName:      "basic",
		StorageName:   "basic",
		MysqlClient:   MysqlClient,
		MongoDBClient: MongoDBClient,
		MinioClient:   MinioClient,
	}
	log.Printf("basic mysql is %p", bk.bookNameMap["basic"].MysqlClient)
}
func GetBookNameInstance() *bookNameMap {
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

// FindBookNameByGorm Deprecated 弃用
func FindBookNameByGorm(db *gorm.DB) (id, name string) {
	log.Printf("find basic mysql is %p", bk.bookNameMap["basic"].MysqlClient)
	GetBookNameInstance().Lock()
	defer GetBookNameInstance().Unlock()
	bk := GetBookNameInstance().bookNameMap
	for key := range bk {
		log.Println(key)
		log.Printf("bk is  %p and create db is %p", bk[key].MysqlClient, db)
		if bk[key].MysqlClient == db {
			return bk[key].StorageName, bk[key].BookName
		}
	}
	return
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
	initBookName()
}

// MysqlScopePaginate 分页函数 纯 gorm 时在 scope 里调用
func MysqlScopePaginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
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

func Paginate(ctx *gin.Context) (int, int) {
	pageNumber := ctx.PostForm("page_number")
	pageSize := ctx.PostForm("page_size")
	log.Println(pageNumber, pageSize)
	var n, limit = 0, 0
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

	return offset, limit
}
