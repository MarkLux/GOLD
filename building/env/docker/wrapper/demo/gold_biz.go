package gold

import (
	"fmt"
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"log"
)

type RedisModel struct {
	Key string
}

func (s *GoldService) Handle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error {
	// the rpc provider of hello service.
	fmt.Printf("Get Request: %v \n", req)
	name := req.Data["name"]
	data := make(map[string]interface{})
	data["rpcResult"] = fmt.Sprintf("Hello, %s", name)
	// the cache for redis
	m := &RedisModel{Key: "Value"}
	err := s.CacheClient.Set("testKey", m, 360 * 1000)
	if err != nil {
		log.Println(err)
	}
	var mm interface{}
	mm, err = s.CacheClient.Get("testKey")
	if err != nil {
		log.Println(err)
	}
	data["cacheResult"] = mm
	rsp.Data = data
	return nil
}