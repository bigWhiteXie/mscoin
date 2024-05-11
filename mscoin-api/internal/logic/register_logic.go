package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"ucenter-api/internal/svc"
)

type UcenterapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUcenterapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UcenterapiLogic {
	return &UcenterapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
