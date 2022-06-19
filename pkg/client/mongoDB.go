package client

import (
	"context"
	cfg "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/qiniu/qmgo"
)

// InitMongoDB ...
func InitMongoDB(config cfg.MongoDBConfig) (*qmgo.Database, error) {

	// 连接到MongoDB
	client, err := qmgo.NewClient(context.TODO(), &qmgo.Config{
		Uri: config.URL,
	})
	if err != nil {
		return nil, err
	}

	return client.Database(config.DBName), nil
}
