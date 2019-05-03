package gold

import "github.com/MarkLux/GOLD/serving/rpc/goldrpc"

// the service function definition
type ServiceFunction interface {
	OnInit()
	OnHandle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error
	OnError(err error) bool
}
