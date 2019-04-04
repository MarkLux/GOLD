# Quick start of GOLD serving runtime

The function services running on GOLD have 3 main capacities:

- RPC: to invoke other function services.
- DB access: to query from the built-in database cluster.
- Cache: set & get data from the built-in cache cluster.

### Usages

- RPC

    You can use the `RpcFactory` of `GoldSerivce` to create a remote service consumer:
    
    ```go
    // create a remote service consumer of 'hello-service' with timeout(3000ms)
    s.RpcFactory.NewRemoteServiceConsumer("hello-service", 3000)
    ```
    
    Thus, an implement of `ServiceConsumer` would be returned, which declared as:
    
    ```go
    type ServiceConsumer interface {
    	Request(request map[string]interface{}) (map[string]interface{}, error)
    }
    ```
    
- DB
 
    GOLD FaaS use NoSQL database as service persist storage (default: MongoDB).

    You can use the `DbFactory` of `GoldService` to create db session or client:
    
    ```go
    // create db session directly
    session := s.DbFactory.NewDataBaseSession("dbName", "tableName", "dbUser", "password")
    // create db client, then create db session
    dbClient := DbFactory.NewDataBaseClient("dbName", "dbUser", "password")
    s1 := dbClient.NewSession("table1")
    s2 := dbClient.NewSession("table2")
    // NOTICE: remember to close the session after you created it:
    defer session.Close()
    defer s1.Close()
    defer s2.Close()
    ```
    
    The session is a connection of a specified table, which is defined as `GoldDataBaseSession`:
    
    ```go
    type GoldDataBaseSession interface {
    	// single data handlers
    	Get(id string) (data GoldDO, err error)
    	Insert(data interface{}) error
    	Update(do GoldDO) error
    	Delete(id string) error
    	// batch data handlers
    	Query(q GoldDBQuery) (data []GoldDO, err error)
    	// close the session and connection.
    	Close()
    }
    ```
    
    Further more, the `GoldDO` is defined as:
    
    ```go
    type GoldDO struct {
    	// unique data id
    	Id string
    	// would be rewritten into json
    	Data interface{}
    	// timestamp of create
    	CreatedAt int64
    	// timestamp of update
    	UpdatedAt int64
    }
    ```
    
    and the `GoldDBQuery` is defined as:
    
    ```go
    type GoldDBQuery struct {
    	Skip  int
    	Limit int
    	Param map[string]string
    }
    ```
    
- Cache

    The `CacheClient` of `GoldService` provide a key-value cache storage for each function service, which is defined as `GoldCacheClient`:
    
    ```go
    type GoldCacheClient interface {
    	Set(key string, val interface{}, expireTime int64) error
    	Get(key string) (interface{}, error)
    }
    ```
    
    **NOTICE: Each function service has its own namespace, which means a same key in different function service has different value**
    
### Full Example

> also see at (serving/wrapper/gold/gold_biz_demo.go)

```go
package gold

import (
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"github.com/MarkLux/GOLD/serving/wrapper/db"
	"log"
)

/**
  * function service example
  * show usage of rpc, db & cache
 */

// the model of user
// use annotation `bson` to control the key saved in db(for mongo).
type UserModel struct {
	Name string `bson:"name"`
	Sex  string `bson:"sex"`
	Mail string `bson:"mail"`
}

// the biz function
func (s *GoldService) Handle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error {
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
		greetingService := s.RpcFactory.NewRemoteServiceConsumer("hello-service", 3000)
		rpcReq := make(map[string]interface{})
		rpcReq["name"] = userName
		greetings, err := greetingService.Request(rpcReq)
		if err != nil {
			log.Println("fail to invoke rpc service, ", err)
		}
		rsp.Data["greetings"] = greetings
	}

	return nil
}

```