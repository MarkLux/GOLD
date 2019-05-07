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
	act := service.Action{
		FunctionService: f,
		TargetBranch: "master",
		TargetVersion: "e37a0171a7a253ecaa57d7b811e11a797d9ba3f4",
		Operator: u,
		Type: "PUBLISH",
	}
	s.PublishFunctionService(act)
}
