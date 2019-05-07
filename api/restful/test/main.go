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
		TargetVersion: "ef13b23aa54102badb2f3d4b6c10067456645b6f",
		Operator: u,
		Type: "PUBLISH",
	}
	s.PublishFunctionService(act)
}
