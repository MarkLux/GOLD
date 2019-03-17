package goldrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MarkLux/GOLD/serving/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type GoldRpcServer struct {
	BindPort string
	ServiceName string
	BizHandler GoldBizHandler
}

// implement the RpcServer using biz handler
func (s *GoldRpcServer) Call(ctx context.Context, req *SyncRequest) (rsp *SyncResponse, err error) {
	// wrap the raw request
	reqData := make(map[string]interface{})
	err = json.Unmarshal(req.Data.Data, &reqData)
	if err != nil {
		return
	}
	goldReq := &GoldRequest{
		Invoker: req.Data.Sender,
		TimeStamp: req.Data.Timestamp,
		Data: reqData,
	}
	goldRsp := &GoldResponse{}
	err = s.BizHandler.Handle(goldReq, goldRsp)
	if err != nil {
		return
	}
	goldRsp.TimeStamp = time.Now().Unix()
	goldRsp.Handler = common.GetGoldEnv().PodName
	// transfer the response
	b, err := json.Marshal(&goldRsp.Data)
	if err != nil {
		return
	}
	d := &SyncData{
		Data: b,
		Sender: goldRsp.Handler,
		Timestamp: goldRsp.TimeStamp,
	}
	rsp = &SyncResponse{Data: d}
	return
}

func (s *GoldRpcServer) Serve() error {
	bindPort := fmt.Sprintf(":%s", s.BindPort)
	lis, err := net.Listen("tcp", bindPort)
	if err != nil {
		return err
	}
	// implement Call using biz handler
	grpcServer := grpc.NewServer()
	RegisterRpcServer(grpcServer, s)
	reflection.Register(grpcServer)
	// bind and service
	if err := grpcServer.Serve(lis); err != nil{
		return err
	}
	log.Printf("serve started succeed.")
	return nil
}