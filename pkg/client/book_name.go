package client

import (
	"github.com/minio/minio-go/v7"
	"github.com/qiniu/qmgo"
	"gorm.io/gorm"
)

// BookName 账套
type BookName struct {
	BookName      string
	StorageName   string
	MysqlClient   *gorm.DB
	MongoDBClient *qmgo.Database
	MinioClient   *minio.Client
}
