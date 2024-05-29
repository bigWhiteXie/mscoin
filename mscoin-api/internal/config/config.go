package config

import (
	"coin-common/queue"
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
	KafkaConfig     queue.KafkaConfig
}

type Mysql struct {
	DataSource string
}

type JWT struct {
	AccessSecret string
	AccessExpire int64
}
