package common

type GoldRequest struct {
	Invoker string
	Data map[string]interface{}
	TimeStamp int64
}

type GoldResponse struct {
	Handler string
	Data map[string]interface{}
	TimeStamp int64
}

// the service function definition
type ServiceFunction interface {
	OnInit()
	OnHandle(req *GoldRequest, rsp *GoldResponse) error
	OnError(err error) bool
}
