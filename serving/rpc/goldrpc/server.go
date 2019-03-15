package goldrpc

import (
	log "log"
	context "context"
	time "time"
	net "net"
	errors "errors"
	grpc "google.golang.org/grpc"
)

type Server struct {}

// rpc implement
// TODO mount the biz function here later
func (*Server) Call(ctx context.Context, req *SyncRequest) (*SyncResponse, error) {
	data := req.Data
	// log the data
	log.Printf("invoker: %s, endpoint: %s, timestamp: %d \n", data.Sender, data.Endpoint, data.Timestamp)
	// return the mock response here.
	resData := &SyncData {
		Sender: "mocked service",
		Endpoint: "127.0.0.1",
		Timestamp: time.Now().Unix(),
	}
	res := &SyncResponse {
		Data: resData,
	}
	return res, nil
}

// lanuch the server
func (*Server) Serve(port string) error {
	// listen the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fatal error happend while listen the port, get error: %v", err)
		return errors.New("fail to listen port.")
	}
	s := grpc.NewServer()
	// bind and register
	RegisterGoldRpcServer(s, &Server{})
	s.Serve(lis)
	return nil
}