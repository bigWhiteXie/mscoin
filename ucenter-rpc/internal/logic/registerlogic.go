package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/register"
	"math/rand"
	"strconv"
	"ucenter/internal/domain"
	"ucenter/internal/svc"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberDomain *domain.MemberDomain
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext, memberDomain *domain.MemberDomain) *RegisterLogic {
	return &RegisterLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		memberDomain: memberDomain,
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	//1.校验手机号验证码是否正确
	key := registerCodeKey(in.Phone)
	if code, err := l.svcCtx.RedisCache.Get(key); err != nil || code != in.Code {
		if err == nil {
			err = fmt.Errorf("验证码错误")
		}
		return nil, err
	}
	l.svcCtx.RedisCache.Del(key)

	//2.查看手机号是否注册
	if mem, err := l.memberDomain.FindByPhone(in.Phone); err != nil || mem.Id != 0 {
		if err == nil {
			err = fmt.Errorf("该手机号已经注册过")
		}
		return nil, err
	}

	//3.注册用户
	if err := l.memberDomain.Register(in.Username, in.Password, in.Phone, in.Country, in.Promotion, in.SuperPartner); err != nil {
		return nil, err
	}
	return &register.RegRes{}, nil
}

func (l *RegisterLogic) SendCode(in *register.CodeReq) (*register.CodeRes, error) {
	//1.生成验证码
	n := rand.Intn(9999)
	if n < 1000 {
		n += 1000
	}
	code := strconv.Itoa(n)
	//2.todo: 手机号发送次数+1,达到当天最大次数则封禁
	l.svcCtx.RedisCache.Setex(registerCodeKey(in.Phone), code, 120)
	//3. 调用api发送短信
	go func() {
		logx.Info("发送短信验证码：" + code)
	}()

	return &register.CodeRes{Code: code}, nil
}

func registerCodeKey(phone string) string {
	return "register:code:" + phone
}
