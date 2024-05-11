package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"market/internal/config"
	"market/internal/database"
)

type ServiceContext struct {
	Config     config.Config
	RedisCache *redis.Redis
	DB         *gorm.DB
	Mongo      *mongo.Database
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	mongoClient := database.ConnectMongo(&c)
	if err != nil {
		panic("数据库连接失败," + c.Mysql.DataSource)
	}
	return &ServiceContext{
		Config:     c,
		RedisCache: redis.MustNewRedis(c.RedisConf, func(r *redis.Redis) {}),
		DB:         db,
		Mongo:      mongoClient.Db,
	}
}
