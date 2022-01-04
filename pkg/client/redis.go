package client

import (
	cfg "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/go-redis/redis/v7"
)

// InitRedis ...
func InitRedis(config cfg.RedisConfig) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.AddrList[0],
		Password: config.Password,
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic("Init Redis failed")
	}
	RedisClient = redisClient
}

func InitRedisCluster(config cfg.RedisConfig) {

	redisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    config.AddrList,
		Password: config.Password,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic("Init redis cluster failed!")
	}
	RedisClusterClient = redisClient
}
