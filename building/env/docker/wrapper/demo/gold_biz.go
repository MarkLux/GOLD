package gold

import (
	"fmt"
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"github.com/MarkLux/GOLD/serving/wrapper/db"
	"log"
)

type RedisModel struct {
	Key string
}

type UserModel struct {
	Name string `bson:"name"`
	Sex string `bson:"sex"`
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
	// the db for mongo
	dbClient := db.NewMongoClient("test", "main_admin", "abc123")
	session, err := dbClient.NewSession("tb")
	if err != nil {
		log.Println("db session err: ", err)
	} else {
		do := db.GoldDO{
			Data: UserModel{
				Name: "lumin",
				Sex: "man",
			},
		}
		err = session.Insert(do)
		if err != nil {
			log.Println("insert err: ", err)
		} else {
			p := make(map[string]string)
			p["name"] = "lumin"
			doList, err := session.Query(db.GoldDBQuery{
				Param: p,
			})
			if err != nil {
				log.Println("query err: ", err)
			} else {
				data["dbResult"] = doList
			}
		}
	}
	rsp.Data = data
	return nil
}