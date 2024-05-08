package logic

import (
	"context"
	"grpc-common/market/rate"

	"github.com/zeromicro/go-zero/core/logx"
	"market/api/internal/svc"
)

type UsdRateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsdRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsdRateLogic {
	return &UsdRateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UsdRateLogic) UsdRate(in *rate.RateReq) (*rate.RateRes, error) {
	// todo: add your logic here and delete this line

	return &rate.RateRes{}, nil
}
