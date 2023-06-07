package svc

import (
	"adminv-api/api/internal/config"
	"adminv-api/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Sys    sysclient.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Sys:    sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
