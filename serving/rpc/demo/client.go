package main

import (
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"log"
	"time"
)

const (
	serviceName  = "demoClient"
	endpoint     = "127.0.0.1"
	providerAddr = "localhost:8080"
)

func main() {
	client := &goldrpc.GoldRpcClient{
		TargetIP:   "127.0.0.1",
		TargetPort: "8080",
		TimeOut:    3000,
	}
	m := make(map[string]interface{})
	m["Name"] = "mark"
	m["Age"] = 22
	req := &goldrpc.GoldRequest{
		Invoker:   serviceName,
		TimeStamp: time.Now().Unix(),
		Data: m,
	}
	rsp, err := client.RequestSync(req)
	if err != nil {
		log.Fatalf("err, %s", err.Error())
	}
	log.Printf("rsp: %v", rsp)
}
