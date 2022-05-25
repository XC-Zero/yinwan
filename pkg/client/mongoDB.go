package client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"

	cfg "github.com/XC-Zero/yinwan/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDB ...
func InitMongoDB(config cfg.MongoDBConfig) (*mongo.Database, error) {

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(config.URL)
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// 检查连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client.Database(config.DBName), nil
}
