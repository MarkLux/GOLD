package cache

import "fmt"

// interface for cache
type GoldCacheClient interface {
	Set(key string, val interface{}, expireTime int64) error
	Get(key string) (interface{}, error)
}

// error
type InvalidParamErr struct {
	Message string
}

func (e *InvalidParamErr) Error() string {
	return e.Message
}

type SerializeErr struct {
	Target interface{}
}

func (e *SerializeErr) Error() string {
	return fmt.Sprintf("fail to serialize %v into json.", e.Target)
}
