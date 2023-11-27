package conn

import (
	"com.ledger.goproject/myconfig"
	"fmt"
	"gopkg.in/redis.v5"
	"strconv"
)

var RClient *redis.Client

func InitRedisConfig() error {
	client := redis.NewClient(&redis.Options{
		Network:  myconfig.GConfig.RedisConfig.Network,
		Addr:     myconfig.GConfig.RedisConfig.IP + ":" + strconv.Itoa(myconfig.GConfig.RedisConfig.Port),
		Password: myconfig.GConfig.RedisConfig.Password,
		DB:       0,
	})
	err := client.Ping().Err()
	if err != nil {
		return fmt.Errorf("redis error: %s\n", err)
	}
	RClient = client
	return nil
}
