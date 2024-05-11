package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/ucenter/login"
	types "ucenter-api/internal/types/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
	"ucenter-api/internal/svc"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	loginRpc := l.svcCtx.LoginRpc
	rpcReq := &login.LoginReq{}
	if err = copier.Copy(rpcReq, req); err != nil {
		return nil, err
	}

	res, err := loginRpc.Login(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	resp = &types.LoginRes{}
	if err = copier.Copy(resp, res); err != nil {
		return nil, err
	}
	return resp, nil
}
