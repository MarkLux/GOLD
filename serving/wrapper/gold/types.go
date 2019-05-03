package gold

import "github.com/MarkLux/GOLD/serving/rpc/goldrpc"

// the service function definition
type ServiceFunction interface {
	onInit()
	onHandle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error
	onError(err error) bool
}
