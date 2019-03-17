package gold

import "github.com/MarkLux/GOLD/serving/wrapper/rpc"

func (s *GoldService) LaunchService() error {
	p := rpc.NewServiceProvider(s)
	return p.Serve()
}
