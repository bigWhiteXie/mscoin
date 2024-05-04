package svc

import (
	"job-center/internal/config"
)

type ServiceContext struct {
	Config *config.Config
}

func NewServiceContext(c *config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config: c,
	}
	return svc
}
