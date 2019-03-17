package goldrpc

import (
	"context"
	"encoding/json"
	"github.com/MarkLux/GOLD/serving/common"
	"google.golang.org/grpc"
	"time"
)

type GoldRpcClient struct {
	TargetIP   string
	TargetPort string
	TimeOut    int64
}

func (client *GoldRpcClient) RequestSync(request *GoldRequest) (response *GoldResponse, err error) {
	response = nil
	// encode the data into json
	jsonBytes, err := json.Marshal(request.Data)
	if err != nil {
		return
	}
	reqData := &SyncData{
		Sender:    common.GetGoldEnv().PodName,
		Data:      jsonBytes,
		Timestamp: time.Now().Unix(),
	}
	// create rpc client
	reqAddr := client.TargetIP + ":" + client.TargetPort
	conn, err := grpc.Dial(reqAddr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return
	}
	c := NewRpcClient(conn)
	// set the timeout in context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(client.TimeOut)*time.Millisecond)
	defer cancel()
	// remote call
	res, err := c.Call(ctx, &SyncRequest{Data: reqData})
	if err != nil {
		return
	}
	// transfer the result into response
	resMap := make(map[string]interface{})
	err = json.Unmarshal(res.Data.Data, &resMap)
	if err != nil {
		return
	}
	response = &GoldResponse{
		Handler:   res.Data.Sender,
		TimeStamp: res.Data.Timestamp,
		Data: resMap,
	}
	return
}
