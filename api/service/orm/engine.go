package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/MarkLux/GOLD/api/service/constant"
	"log"
	"sync"
)

// singleton engine instance here.
var engine *xorm.Engine

var once sync.Once

func InitOrmEngine() {
	once.Do(func() {
		dialUrl := "root:qwe123@localhost:3306/gold?charset=UTF-8"
		var err error
		engine, err = xorm.NewEngine(constant.DataBaseDriver, dialUrl)
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