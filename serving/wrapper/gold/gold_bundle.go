package gold

import (
	"fmt"
	"github.com/MarkLux/GOLD/serving/wrapper/cache"
)

type LaunchError struct {
	Message string
}

func (e *LaunchError) Error() string {
	return e.Message
}

type GoldService struct {
	// singleton component, inject here.
	CacheClient cache.GoldCacheClient
}

func (s *GoldService) LoadComponents() error {
	// load all the required injection here.
	var err error
	s.CacheClient, err = cache.GetGoldRedisClient()
	if err != nil {
		// init failed.
		return &LaunchError{Message: fmt.Sprintf("fail to init redis client, %s", err.Error())}
	}
	return nil
}
