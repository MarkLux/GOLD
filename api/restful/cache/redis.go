package cache

import (
	"github.com/MarkLux/GOLD/api/restful/constant"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

type RedisClient struct {
	rClient *redis.Client
}

// singleton
var instance *RedisClient
var once sync.Once

func GetRedisClient() *RedisClient {
	if instance != nil {
		return instance
	}
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr: constant.RedisAddr,
		})
		instance = &RedisClient{
			rClient: client,
		}
	})
	return instance
}

func (c RedisClient) Set(key string, value string, expireTime int64) error {
	return c.rClient.Set(key, value, time.Duration(expireTime) * time.Millisecond).Err()
}

func (c RedisClient) Get(key string) (string, error) {
	return c.rClient.Get(key).Result()
}