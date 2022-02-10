package client

import (
	"github.com/XC-Zero/yinwan/pkg/config"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// BookName 账套
type BookName struct {
	Name          string
	MysqlClient   *gorm.DB
	MongoDBClient *mongo.Database
	MinioClient   *minio.Client
}

func InitBookMap(configs []config.BookConfig) {
	for _, config := range configs {
		minio, err := InitMinio(config.MinioConfig)
		if err != nil {
			//	todo error list
		}
		bk := BookName{
			MysqlClient:   InitMysqlGormV2(config.MysqlConfig),
			Name:          config.BookName,
			MongoDBClient: InitMongoDB(config.MongoDBConfig),
			MinioClient:   minio,
		}
		BookNameMap[config.BookName] = bk
	}

}
