package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"market/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	RedisCache *redis.Redis
	DB         *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("数据库连接失败," + c.Mysql.DataSource)
	}
	return &ServiceContext{
		Config:     c,
		RedisCache: redis.MustNewRedis(c.RedisConf, func(r *redis.Redis) {}),
		DB:         db,
	}
}
