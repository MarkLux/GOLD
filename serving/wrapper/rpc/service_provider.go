package rpc

import (
	"github.com/MarkLux/GOLD/serving/common"
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"github.com/MarkLux/GOLD/serving/wrapper/constant"
)

// interface
type ServiceProvider interface {
	Serve() error
}

type GoldServiceProvider struct {
	server goldrpc.GoldRpcServer
	Handler goldrpc.GoldBizHandler
}

func (p *GoldServiceProvider) Serve() error {
	p.server = goldrpc.GoldRpcServer{
		BindPort: constant.DefaultServicePort,
		ServiceName: common.GetGoldEnv().ServiceName,
		BizHandler: p.Handler,
	}
	return p.server.Serve()
}

func NewServiceProvider(h goldrpc.GoldBizHandler) *GoldServiceProvider {
	return &GoldServiceProvider{Handler: h}
}