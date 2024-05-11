package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf redis.RedisConf
	Mysql     Mysql
	JWT       JWT
	Mongo     Mongo
}
type Mongo struct {
	Url string
}

type Mysql struct {
	DataSource string
}

type JWT struct {
	AccessSecret string
	AccessExpire int64
}
