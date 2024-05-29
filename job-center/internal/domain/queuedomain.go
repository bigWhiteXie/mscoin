package domain

import (
	"coin-common/queue"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"job-center/internal/model"
	"job-center/internal/svc"
)

type QueueDomain struct {
	kafkaClient *queue.KafkaClient
}

const KLINE_1M_TOPIC = "kline_1m"
const KLINE_1H_TOPIC = "kline_1h"

func NewQueueDomain(svr *svc.ServiceContext) *QueueDomain {
	return &QueueDomain{
		kafkaClient: svr.Kafka,
	}
}

func (q *QueueDomain) PushKline(data []string, symbol string, period string) error {
	topic := KLINE_1M_TOPIC
	if period == "1h" {
		topic = KLINE_1H_TOPIC
	}
	kafkaData := &queue.KafkaData{
		Topic: topic,
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
