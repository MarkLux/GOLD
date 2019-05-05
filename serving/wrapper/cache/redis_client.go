package cache

import (
	"encoding/json"
	"github.com/MarkLux/GOLD/serving/common"
	"github.com/go-redis/redis"
	"log"
	"sync"
	"time"
)

const DefaultPrefix = "gold_redis_key_"

var instance *GoldRedisClient
var once sync.Once

func GetGoldRedisClient() (ins *GoldRedisClient, err error) {
	once.Do(func() {
		instance, err = NewGoldRedisClient()
	})
	ins = instance
	return
}

type GoldRedisClient struct {
	rClient *redis.ClusterClient
}

func NewGoldRedisClient() (*GoldRedisClient, error) {
	c := redis.NewClusterClient(&redis.ClusterOptions{
		// why can't I use a single restful name?
		// fix: do not use hard codec
		Addrs: []string {
			"redis-app-0.redis-restful.gold.svc.cluster.local:6379",
			"redis-app-1.redis-restful.gold.svc.cluster.local:6379",
			"redis-app-2.redis-restful.gold.svc.cluster.local:6379",
			"redis-app-3.redis-restful.gold.svc.cluster.local:6379",
			"redis-app-4.redis-restful.gold.svc.cluster.local:6379",
			"redis-app-5.redis-restful.gold.svc.cluster.local:6379",
		},
	})
	// test conn
	pong, err := c.Ping().Result()
	if err != nil {
		log.Printf("fail to init redis client, %s", err.Error())
		return nil, err
	}
	log.Printf("succeed connect to redis cluster, server pong: %s", pong)
	return &GoldRedisClient{rClient: c}, nil
}

func (r *GoldRedisClient) Set(key string, val interface{}, expireTime int64) error {
	if key == "" {
		return &InvalidParamErr{Message: "cache key should not be empty!"}
	}
	k := genKey(key)
	// transfer value into json
	v, err := json.Marshal(val)
	if err != nil {
		return &SerializeErr{Target: val}
	}
	return r.rClient.Set(k, v, time.Duration(expireTime)*time.Millisecond).Err()
}

func (r *GoldRedisClient) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, &InvalidParamErr{Message: "cache key should not be empty!"}
	}
	k := genKey(key)
	res, err := r.rClient.Get(k).Result()
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal([]byte(res), &v)
	if err != nil {
		return nil, &SerializeErr{Target: res}
	}
	return v, nil
}

func genKey(raw string) string {
	return DefaultPrefix + common.GetGoldEnv().ServiceName + "_" + raw
}
