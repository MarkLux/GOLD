package goldrpc

type GoldRequest struct {
	Invoker string
	Data interface{}
	TimeStamp int64
}

type GoldResponse struct {
	Handler string
	Data interface{}
	TimeStamp int64
}
