package svc

import (
	"adminv-api/api/internal/config"
	"adminv-api/api/internal/middleware"
	"adminv-api/rpc/sys/sys_client"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	Sys      sys_client.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.MustNewRedis(redisConfig(c))
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware(newRedis).Handle,
		Sys:      sys_client.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}

func redisConfig(c config.Config) redis.RedisConf {
	return redis.RedisConf{
		Host:        c.Redis.Host,
		Type:        redis.NodeType,
		Pass:        c.Redis.Pass,
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 0,
	}
}
