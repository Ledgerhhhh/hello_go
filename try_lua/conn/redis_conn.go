package conn

import (
	"github.com/go-redis/redis/v8"
)

var RClient *redis.Client

func InitRedisConfig() error {
	RClient = redis.NewClient(&redis.Options{
		Addr:     "106.54.9.19:6379", // Redis 服务器地址
		Password: "biaoge666",        // Redis 密码，如果有的话
		DB:       0,                  // 使用默认数据库
	})
	return nil
}
