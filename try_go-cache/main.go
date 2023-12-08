package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	// 设置5分钟过期,10分钟清理
	c := cache.New(5*time.Minute, 10*time.Minute)
	// 默认的时间5分钟
	c.Set("key", "value", cache.DefaultExpiration)
	// 不过期
	c.Set("key2", "value2", cache.NoExpiration)
	// 12秒过期
	c.Set("key3", "value3", 12*time.Second)
	c.Delete("key3")
	for key, item := range c.Items() {
		fmt.Printf("Key: %s, Expiration: %d\n", key, item.Expiration)
	}
}
