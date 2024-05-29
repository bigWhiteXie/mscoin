package processor

import (
	"coin-common/queue"
	"encoding/json"
	"grpc-common/market"
	"ucenter-api/internal/model"
)

const (
	KLINE = "kline"
)

type KlineProcessor struct {
	thumbMap map[string]*market.CoinThumb
	handlers []MarketHandler
	kafkaCli *queue.KafkaClient
}

func NewKlineProcessor(kafkaCli *queue.KafkaClient) *KlineProcessor {
	return &KlineProcessor{
		kafkaCli: kafkaCli,
		handlers: make([]MarketHandler, 0),
		thumbMap: make(map[string]*market.CoinThumb),
	}
}

func (d *KlineProcessor) Process(data ProcessData) {
	if data.Type == KLINE {
		symbol := string(data.Key)
		kline := &model.Kline{}
		json.Unmarshal(data.Data, kline)
		for _, v := range d.handlers {
			v.HandleKLine(symbol, kline, d.thumbMap)
		}
	}
}

func (d *KlineProcessor) AddHandler(h MarketHandler) {
	//发送到websocket的服务
	d.handlers = append(d.handlers, h)
}

func (k *KlineProcessor) GetThumb() any {
	return k.thumbMap
}
