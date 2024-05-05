package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/ucenter/login"
	"grpc-common/ucenter/register"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RegisterRpc register.RegisterClient
	LoginRpc    login.LoginClient
	RedisCache  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RegisterRpc: register.NewRegisterClient(zrpc.MustNewClient(c.Rpcclient).Conn()),
		LoginRpc:    login.NewLoginClient(zrpc.MustNewClient(c.Rpcclient).Conn()),
		RedisCache:  redis.MustNewRedis(c.Redis, func(r *redis.Redis) {}),
	}
}
