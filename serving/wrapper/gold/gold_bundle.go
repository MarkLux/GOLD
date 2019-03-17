package gold

import (
	"github.com/MarkLux/GOLD/serving/wrapper/rpc"
)

type GoldService struct {
	RpcConsumer rpc.GoldServiceConsumer
}

func (s *GoldService) LoadComponents() {
	// load all the required injection here.
}
