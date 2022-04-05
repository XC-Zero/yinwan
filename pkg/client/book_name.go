package client

import (
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
