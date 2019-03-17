package goldrpc

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

type GoldBizHandler interface {
	Handle(request *GoldRequest, response *GoldResponse) error
}