// Code generated by goctl. DO NOT EDIT.
// Source: login.proto

package server

import (
	"context"
	"grpc-common/ucenter/login"
	"ucenter/internal/domain"
	"ucenter/internal/logic"
	"ucenter/internal/svc"
)

type LoginServer struct {
	svcCtx *svc.ServiceContext
	login.UnimplementedLoginServer
}

func NewLoginServer(svcCtx *svc.ServiceContext) *LoginServer {
	return &LoginServer{
		svcCtx: svcCtx,
	}
}

func (s *LoginServer) Login(ctx context.Context, in *login.LoginReq) (*login.LoginRes, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx, domain.NewMemberDomain(s.svcCtx.DB))
	return l.Login(in)
}
