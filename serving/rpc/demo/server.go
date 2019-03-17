package main

import (
	"fmt"
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"github.com/MarkLux/GOLD/serving/wrapper/constant"
	"log"
)

type bizFunc struct {}

func (*bizFunc) Handle(request *goldrpc.GoldRequest, response *goldrpc.GoldResponse) error {
	log.Printf("got request: %v", request)
	name := request.Data["Name"]
	age := request.Data["Age"]
	m := make(map[string]interface{})
	m["Result"] = fmt.Sprintf("hello %s, your age is %f", name, age)
	response.Data = m
	fmt.Println(m)
	return nil
}

func main() {
	server := &goldrpc.GoldRpcServer{
		BindPort: constant.DefaultServicePort,
		ServiceName: "helloService",
		BizHandler: &bizFunc{},
	}
	err := server.Serve()
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}
}
