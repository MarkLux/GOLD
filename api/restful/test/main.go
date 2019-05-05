package main

import (
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/MarkLux/GOLD/api/restful/service"
)

func main() {
	s := service.GetFunctionService()
	f := s.GetFunctionService(int64(2))
	u := orm.User{
		Id: 1,
		Name: "marklux",
	}
	act := service.UpdateAction{
		FunctionService: f,
		TargetBranch: "master",
		TargetVersion: "",
		Operator: u,
	}
	s.PublishFunctionService(act)
}
