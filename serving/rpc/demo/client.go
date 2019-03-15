package main

import (
	"log"
    "os"
  
    "golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
	"goldrpc"
)

const (
	serviceName = "demoClient"
	endpoint = "127.0.0.1"
	providerAddr = "localhost:8099"
)

func main() {
	conn, err := grpc.Dail(providerAddr, grpc.WithInsecure())
	if err != nil {
        log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	
	c := goldrpc.NewGoldRpcClient(conn)

	data := &goldrpc.SyncData{
		Sender: serviceName,
		Endpoint: endpoint,
		Timestamp: time.Now().Unix()
	}

	req := &goldrpc.SyncRequest{
		Data: data
	}

	res, err := c.Call(context.Background(), req)

	if err != nil {
		log.Fatal("fail to get response: %v", err)
		return
	}

	log.Printf("response: %v\n", res.Data)
}