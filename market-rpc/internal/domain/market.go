package domain

import (
	"coin-common/tools"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market"
	"market/internal/dao"
	"market/internal/model"
	"market/internal/svc"
	"math"
	"time"
)

type MarketDomain struct {
	*dao.KlineDao
}

func NewMarketDomain(svr *svc.ServiceContext) *MarketDomain {
	return &MarketDomain{
		dao.NewKlineDao(svr.Mongo),
	}
}

func (d *MarketDomain) SymbolThumbTrend(coin *model.ExchangeCoin) (*market.CoinThumb, error) {
	from := tools.ZeroTime()
	to := time.Now().UnixMilli()
	klines, err := d.FindKlineInTime(coin.Symbol, "1m", from, to)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	if len(klines) <= 0 {
		return &market.CoinThumb{
			Symbol: coin.Symbol,
		}, nil
	}
	trend := make([]float64, len(klines))
	var high float64 = 0
	var low float64 = math.MaxFloat64
	var volumes float64 = 0
	var turnover float64 = 0
	for i, v := range klines {
		trend[i] = v.ClosePrice
		if v.HighestPrice > high {
			high = v.HighestPrice
		}
		if v.LowestPrice < low {
			low = v.LowestPrice
		}
		volumes += v.Volume
		turnover += v.Turnover
	}
	kline := klines[0]
	end := klines[len(klines)-1]
	ct := kline.ToCoinThumb(coin.Symbol, end)
	ct.High = high
	ct.Low = low
	ct.Volume = volumes
	ct.Turnover = turnover
	ct.Trend = trend
	return ct, nil
}
