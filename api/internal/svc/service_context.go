package svc

import (
	"adminv-api/api/internal/config"
	"adminv-api/rpc/sys/sys_client"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Sys    sys_client.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Sys:    sys_client.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
