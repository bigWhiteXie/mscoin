package domain

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"job-center/internal/database"
	"job-center/internal/model"
	"job-center/internal/svc"
)

type QueueDomain struct {
	kafkaClient *database.KafkaClient
}

const KLINE_1M_TOPIC = "kline_1m"

func NewQueueDomain(svr *svc.ServiceContext) *QueueDomain {
	return &QueueDomain{
		kafkaClient: svr.Kafka,
	}
}

func (q *QueueDomain) PushKline(data []string, symbol string, period string) error {
	kafkaData := &database.KafkaData{
		Topic: KLINE_1M_TOPIC,
		Key:   []byte(symbol)}
	kline := model.NewKline(data, period)
	val, err := json.Marshal(kline)
	if err != nil {
		logx.Info("queueDomain: 序列化kline异常")
		return err
	}
	kafkaData.Data = val
	q.kafkaClient.Send(*kafkaData)
	return nil
}
