package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market"
	"time"
	"ucenter-api/internal/svc"
	types "ucenter-api/internal/types/market"
)

type MarketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MarketLogic) GetCoinsThumbTrend(t *types.MarketReq) ([]types.CoinThumbResp, error) {
	ctx, cancelFunc := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancelFunc()
	req := &market.MarketReq{Ip: t.Ip}
	grpcResp, err := l.svcCtx.MarketClient.FindSymbolThumbTrend(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := []types.CoinThumbResp{}
	copier.Copy(&resp, grpcResp.List)
	return resp, nil
}
