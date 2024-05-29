package svc

import (
	"coin-common/queue"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market"
	"grpc-common/market/rate"
	"grpc-common/ucenter/login"
	"grpc-common/ucenter/register"
	"ucenter-api/internal/config"
	"ucenter-api/internal/processor"
)

type ServiceContext struct {
	Config             config.Config
	RegisterRpc        register.RegisterClient
	ExchangeRateClient rate.ExchangeRateClient
	MarketClient       market.MarketClient
	LoginRpc           login.LoginClient
	RedisCache         *redis.Redis
	KafkaClient        *queue.KafkaClient
	KlineProcssor      *processor.KlineProcessor
}

func NewServiceContext(c config.Config) *ServiceContext {
	kafkaClient := queue.NewKafkaClient(&c.KafkaConfig)
	kafkaClient.StartWrite()
	return &ServiceContext{
		Config:             c,
		RegisterRpc:        register.NewRegisterClient(zrpc.MustNewClient(c.UcRpcClient).Conn()),
		ExchangeRateClient: rate.NewExchangeRateClient(zrpc.MustNewClient(c.MarketRpcClient).Conn()),
		LoginRpc:           login.NewLoginClient(zrpc.MustNewClient(c.UcRpcClient).Conn()),
		RedisCache:         redis.MustNewRedis(c.Redis, func(r *redis.Redis) {}),
		MarketClient:       market.NewMarketClient(zrpc.MustNewClient(c.MarketRpcClient).Conn()),
		KafkaClient:        kafkaClient,
	}
}
