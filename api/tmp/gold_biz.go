package gold

import (
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"log"
)

func (s *GoldService) OnInit() {
	log.Println("inited")
}

// the biz function
func (s *GoldService) OnHandle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error {
	// get data from request
	userName := req.Data["name"].(string)
	log.Println("userName: " + userName)

	greeting := "hello, " + userName
	rsp.Data["greeting"] = greeting

	return nil
}

func (s *GoldService) OnError(err error) bool {
	log.Println("error!", err)
	return false
}