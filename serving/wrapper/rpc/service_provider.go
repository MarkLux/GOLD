package rpc

import (
	"github.com/MarkLux/GOLD/serving/common"
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"github.com/MarkLux/GOLD/serving/wrapper/constant"
	"github.com/MarkLux/GOLD/serving/wrapper/gold"
)

// interface
type ServiceProvider interface {
	Serve() error
}

type GoldServiceProvider struct {
	server goldrpc.GoldRpcServer
	Function gold.ServiceFunction
}

func (p *GoldServiceProvider) Serve() error {
	p.server = goldrpc.GoldRpcServer{
		BindPort: constant.DefaultServicePort,
		ServiceName: common.GetGoldEnv().ServiceName,
		Function: p.Function,
	}
	return p.server.Serve()
}

func NewServiceProvider(f gold.ServiceFunction) *GoldServiceProvider {
	return &GoldServiceProvider{Function: f}
}