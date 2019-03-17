package gold

import (
	"fmt"
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"github.com/MarkLux/GOLD/serving/wrapper/rpc"
)

func (s *GoldService) Handle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error {
	// rpc example
	helloService := rpc.GetRemoteService("helloService")
	helloReq := make(map[string]interface{})
	data := make(map[string]interface{})
	helloReq["name"] = "lumin"
	res, err := helloService.Request(helloReq)
	if err != nil {
		data["result"] = fmt.Sprintf("Get Error: %s", err.Error())
	} else {
		data["result"] = fmt.Sprintf("Get Response: %s", res["result"])
	}
	rsp.Data = data
	return nil
}