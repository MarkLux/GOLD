package gold

import (
	"fmt"
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
)

func (s *GoldService) Handle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error {
	// the rpc provider of hello service.
	fmt.Printf("Get Request: %v \n", req)
	name := req.Data["name"]
	data := make(map[string]interface{})
	data["result"] = fmt.Sprintf("Hello, %s", name)
	rsp.Data = data
	return nil
}