package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/rate"
	domain "market/internal/domain/exchange"
	"market/internal/svc"
)

type ExchangeRateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeRateDomain *domain.ExchangeRateDomain
}

func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext, exchangeRateDomain *domain.ExchangeRateDomain) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeRateDomain: exchangeRateDomain,
	}
}

func (l *ExchangeRateLogic) GetUsdRate(in *rate.RateReq) (*rate.RateRes, error) {
	usdRate := l.exchangeRateDomain.GetUsdRate(in.GetUnit())
	res := &rate.RateRes{Rate: usdRate}
	return res, nil
}
