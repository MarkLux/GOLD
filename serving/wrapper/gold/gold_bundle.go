package gold

import (
	"fmt"
	"github.com/MarkLux/GOLD/serving/wrapper/cache"
	"github.com/MarkLux/GOLD/serving/wrapper/db"
	"github.com/MarkLux/GOLD/serving/wrapper/rpc"
)

type LaunchError struct {
	Message string
}

func (e *LaunchError) Error() string {
	return e.Message
}

// rpc factory client
type GoldRpcFactory struct {}

// db factory client
type GoldDbFactory struct {
	Driver string
}

func (GoldRpcFactory) NewRemoteServiceConsumer(serviceName string, timeOut int64) *rpc.GoldServiceConsumer {
	return rpc.GetRemoteServiceWithTimeOut(serviceName, timeOut)
}

func (f GoldDbFactory) NewDataBaseSession(dataBase string, tb string, user string, pwd string) (db.GoldDataBaseSession, error) {
	if f.Driver == "mongo" {
		// currently just bind mongo client here
		return db.NewMongoClient(dataBase, user, pwd).NewSession(tb)
	} else {
		return nil, db.DBCommonError{Message:"Invalid Driver"}
	}
}

func (f GoldDbFactory) NewDataBaseClient(dataBase string, user string, pwd string) db.GoldDataBaseClient {
	if f.Driver == "mongo" {
		// currently just bind mongo client here
		client := db.NewMongoClient(dataBase, user, pwd)
		return &client
	} else {
		return nil
	}
}

type GoldService struct {
	// singleton component, inject here.
	CacheClient cache.GoldCacheClient
	// multi instances use factory.
	DbFactory GoldDbFactory
	RpcFactory GoldRpcFactory
}

func (s *GoldService) LoadComponents() error {
	// create factory
	s.DbFactory = GoldDbFactory{Driver: "mongo"}
	s.RpcFactory = GoldRpcFactory{}
	// create singleton injection
	var err error
	s.CacheClient, err = cache.GetGoldRedisClient()
	if err != nil {
		// init failed.
		return &LaunchError{Message: fmt.Sprintf("fail to init redis client, %s", err.Error())}
	}
	return nil
}
