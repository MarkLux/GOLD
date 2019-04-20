package main

import (
	"github.com/MarkLux/GOLD/api/service/orm"
)

func main() {
	engine := orm.GetOrmEngine()
	if engine == nil {
		panic("init orm engine failed!")
	}
}
