package gold

import (
	"github.com/MarkLux/GOLD/serving/wrapper/rpc"
	"github.com/MarkLux/GOLD/serving/wrapper/cache"
)

type LaunchError struct {
	Message string
}

func (e *LaunchError) Error() string {
	return e.Message
}

type GoldService struct {
	RpcConsumer rpc.GoldServiceConsumer
	CacheClient *cache.GoldCacheClient
}

func (s *GoldService) LoadComponents() error {
	// load all the required injection here.
	s.CacheClient = cache.GetGoldRedisClient()
	if s.CacheClient == nil {
		// init failed.
		return &LaunchError{Message: "fail to init redis client."}
	}
	return nil
}
