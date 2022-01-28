package client

import (
	"context"

	cfg "github.com/XC-Zero/yinwan/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDB ...
func InitMongoDB(config cfg.MongoDBConfig) *mongo.Database {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(config.URL)
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	return client.Database(config.DBName)
}
