package logic

import (
	"context"
	"grpc-common/market"
	"market/internal/domain"
	"market/internal/model"
	"market/internal/svc"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	marketDomain *domain.MarketDomain
	coinDomain   *domain.ExchangeCoinDomain
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		marketDomain: domain.NewMarketDomain(svcCtx),
		coinDomain:   domain.NewExchangeCoinDomain(svcCtx.DB),
	}
}

func (l *MarketLogic) FindSymbolThumbTrend(in *market.MarketReq) (*market.SymbolThumbRes, error) {
	//1.找出可见的所有symbol
	coins := l.coinDomain.FindVisible(l.ctx)
	thumbs := make([]*market.CoinThumb, len(coins))
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(coins))

	//2.遍历coins，得到每个symbol的缩略信息
	for i, coin := range coins {
		go func(coin *model.ExchangeCoin, index int) {
			thumb, err := l.marketDomain.SymbolThumbTrend(coin)
			if err != nil {
				logx.Error(err)
			}
			thumbs[index] = thumb
			waitGroup.Done()
		}(coin, i)
	}
	waitGroup.Wait()
	return &market.SymbolThumbRes{List: thumbs}, nil
}
