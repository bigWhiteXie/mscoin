package logic

import (
	"coin-common/tools"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"grpc-common/ucenter/login"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberDomain *domain.MemberDomain
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, memberDomain *domain.MemberDomain) *LoginLogic {
	return &LoginLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		memberDomain: memberDomain,
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginRes, error) {
	mem, err := l.memberDomain.FindByUsername(in.Username)
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, err
	}
	if mem == nil {
		return nil, errors.New("用户名不存在")
	}
	if ok := tools.Verify(in.Password, mem.Salt, mem.Password, nil); !ok {
		return nil, errors.New("密码不正确")
	}
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret
	token, err := l.getJwtToken(accessSecret, time.Now().Unix(), accessExpire, mem.Id)
	if err != nil {
		return nil, errors.New("未知错误，请联系管理员")
	}
	loginCount := mem.LoginCount + 1
	go func() {
		l.memberDomain.UpdateLoginCount(mem.Id, 1)
	}()
	return &login.LoginRes{
		Token:         token,
		Id:            mem.Id,
		Username:      mem.Username,
		MemberLevel:   mem.MemberLevelStr(),
		MemberRate:    mem.MemberRate(),
		RealName:      mem.RealName,
		Country:       mem.Country,
		Avatar:        mem.Avatar,
		PromotionCode: mem.PromotionCode,
		SuperPartner:  mem.SuperPartner,
		LoginCount:    int32(loginCount),
	}, nil

}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
