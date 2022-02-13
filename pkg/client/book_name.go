package client

import (
	"github.com/XC-Zero/yinwan/pkg/config"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// BookName 账套
type BookName struct {
	BookName      string
	StorageName   string
	MysqlClient   *gorm.DB
	MongoDBClient *mongo.Database
	MinioClient   *minio.Client
}

func InitBookMap(configs []config.BookConfig) error {
	var errorList []error
	for _, config := range configs {
		mi, err := InitMinio(config.MinioConfig)
		if err != nil {
			errorList = append(errorList, err)
		}
		mydb, err := InitMysqlGormV2(config.MysqlConfig)
		if err != nil {
			errorList = append(errorList, err)
		}
		db, err := InitMongoDB(config.MongoDBConfig)
		if err != nil {
			errorList = append(errorList, err)
		}
		bk := BookName{
			MysqlClient:   mydb,
			StorageName:   config.StorageName,
			BookName:      config.BookName,
			MongoDBClient: db,
			MinioClient:   mi,
		}
		BookNameMap[config.BookName] = bk
	}
	return errs.ErrorListToError(errorList)
}
