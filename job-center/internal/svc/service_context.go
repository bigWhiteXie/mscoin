package svc

import (
	"job-center/internal/config"
	"job-center/internal/database"
)

type ServiceContext struct {
	Config *config.Config
	Kafka  *database.KafkaClient
}

func NewServiceContext(c *config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config: c,
		Kafka:  database.NewKafkaClient(&c.Kafka),
	}
	return svc
}
