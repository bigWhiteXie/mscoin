package consumer

import "ucenter-api/internal/svc"

func RegisterConsumer(svrCtx *svc.ServiceContext) {
	client := svrCtx.KafkaClient
	klineConsumer := NewKlineConsumer(svrCtx.KlineProcssor)
	client.RegisterConsumer(klineConsumer)
	client.StartRead()
}
