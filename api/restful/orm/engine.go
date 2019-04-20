package orm

import (
	"github.com/MarkLux/GOLD/api/restful/constant"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

// singleton engine instance here.
var engine *xorm.Engine

var once sync.Once

func InitOrmEngine() {
	once.Do(func() {
		// for mysql only
		dsn := "root:qwe123@tcp(127.0.0.1:3306)/gold?charset=utf8"
		var err error
		engine, err = xorm.NewEngine(constant.DataBaseDriver, dsn)
		if err != nil {
			log.Printf("fail to init orm engine, %s", err.Error())
		}
	})
}

func GetOrmEngine() *xorm.Engine  {
	if engine == nil {
		InitOrmEngine()
	}
	return engine
}