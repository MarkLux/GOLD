package goldrpc

type GoldRpcClient struct {
	TargetIP   string
	TargetPort string
	TimeOut    int32
}

func (*GoldRpcClient) RequestSync(request GoldRequest) (response *GoldResponse, err error) {
	if request != nil {

	}
}
