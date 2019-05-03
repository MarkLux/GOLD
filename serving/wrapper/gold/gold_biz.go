package gold

import (
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"github.com/MarkLux/GOLD/serving/wrapper/db"
	"log"
)

/**
  * function restful example
  * show usage of rpc, db & cache
 */

// the model of user
// use annotation `bson` to control the key saved in db(for mongo).
type UserModel struct {
	Name string `bson:"name"`
	Sex  string `bson:"sex"`
	Mail string `bson:"mail"`
}

func (s *GoldService) OnInit() {
	// do something here..
}

// the biz function
func (s *GoldService) OnHandle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error {
	// get data from request
	userName := req.Data["name"].(string)

	// cache example
	cacheKey := "prefix_" + userName
	u, err := s.CacheClient.Get(cacheKey)
	if err != nil {
		log.Println("fail to get info from cache, ", err)
		return err
	}

	// if got nothing from cache, then query the db.
	if u == nil {
		// db session example
		dbSession, err := s.DbFactory.NewDataBaseSession("test", "user", "root", "pwd")
		if err != nil {
			log.Println("create db session failed, ", err)
			return err
		}
		defer dbSession.Close()
		// db query example
		param := make(map[string]string)
		param["data.name"] = userName
		qUsers, err := dbSession.Query(db.GoldDBQuery{Param: param})
		if err != nil {
			log.Println("fail to query db, ", err)
			return err
		}
		if len(qUsers) > 1 {
			u = qUsers[0]
			// reset the cache
			err = s.CacheClient.Set(cacheKey, u, 300 * 1000)
			if err != nil {
				log.Println("fail to reset cache, ", err)
			}
		}
		u = nil
	}

	// build response
	rsp.Data = make(map[string]interface{})
	if u != nil {
		rsp.Data["userModel"] = u
		// rpc example
		greetingService := s.RpcFactory.NewRemoteServiceConsumer("hello-restful", 3000)
		rpcReq := make(map[string]interface{})
		rpcReq["name"] = userName
		greetings, err := greetingService.Request(rpcReq)
		if err != nil {
			log.Println("fail to invoke rpc restful, ", err)
		}
		rsp.Data["greetings"] = greetings
	}

	return nil
}

func (s *GoldService) OnError(err error) bool {
	log.Println(err.Error())
	return false
}