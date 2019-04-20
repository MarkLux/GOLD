package main

import (
	"github.com/MarkLux/GOLD/api/restful/orm"
	"log"
	"time"
)

func main() {
	engine := orm.GetOrmEngine()
	if engine == nil {
		panic("init orm engine failed!")
	}
	current := time.Now().Unix()
	u := &orm.User{
		Name: "lumin",
		Email: "marlx6590@163.com",
		CreatedAt: current,
		UpdatedAt: current,
		AddOn: "",
	}
	engine.ShowSQL(true)
	affected, err := engine.Insert(u)
	if err != nil {
		panic(err)
	}
	log.Println(string(affected) + "rows inserted.")
}
