package consumer

import (
	"coin-common/queue"
	"ucenter-api/internal/processor"
)

type KlineConsumer struct {
	processor processor.Processor
}

func NewKlineConsumer(p *processor.KlineProcessor) *KlineConsumer {
	return &KlineConsumer{p}
}
func (k *KlineConsumer) Consume(data queue.KafkaData) error {
	pData := processor.ProcessData{
		Type: "kline",
		Key:  data.Key,
		Data: data.Data,
	}
	k.processor.Process(pData)
	return nil
}

func (k *KlineConsumer) Topic() string {
	return "kline"
}

func (k *KlineConsumer) Group() string {
	return "kline"
}
