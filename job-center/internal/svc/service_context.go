package svc

import (
	"coin-common/queue"
	"job-center/internal/config"
)

type ServiceContext struct {
	Config *config.Config
	Kafka  *queue.KafkaClient
}

func NewServiceContext(c *config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config: c,
		Kafka:  queue.NewKafkaClient(&c.Kafka),
	}
	svc.Kafka.StartWrite()
	return svc
}
