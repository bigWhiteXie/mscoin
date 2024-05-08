package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UcRpcClient     zrpc.RpcClientConf
	MarketRpcClient zrpc.RpcClientConf
	Mysql           Mysql
	Redis           redis.RedisConf
	Jwt             JWT
}

type Mysql struct {
	DataSource string
}

type JWT struct {
	AccessSecret string
	AccessExpire int64
}
