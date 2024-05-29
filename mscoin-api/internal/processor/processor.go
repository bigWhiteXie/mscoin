package processor

import (
	"grpc-common/market"
	"ucenter-api/internal/model"
)

type MarketHandler interface {
	HandleTrade(symbol string, data []byte)
	HandleKLine(symbol string, kline *model.Kline, thumbMap map[string]*market.CoinThumb)
}

type Processor interface {
	GetThumb() any
	Process(data ProcessData)
	AddHandler(h MarketHandler)
}

type ProcessData struct {
	Type string //trade 交易 kline k线
	Key  []byte
	Data []byte
}
