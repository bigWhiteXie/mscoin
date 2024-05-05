package handler

import (
	"coin-common/common"
	"coin-common/validate"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"grpc-common/ucenter/register"
	"net/http"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type RegisterHandler struct {
	svcCtx *svc.ServiceContext
}

const (
	PHONE_CODE_PREFIX = "code:phone:"
	CODE_EXPIRE       = 120
)

func NewRegisterHandler(ctx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{svcCtx: ctx}
}

func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req *register.RegReq
	if err := httpx.Parse(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	_, err := h.svcCtx.RegisterRpc.RegisterByPhone(ctx, req)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, common.NewResult().Success("ok"))
	}
}

func (h *RegisterHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req *types.CodeReq
	if err := httpx.Parse(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	//1.校验手机号格式是否正确
	if ok := validate.IsPhoneFomatter(req.Phone); !ok {
		httpx.OkJsonCtx(r.Context(), w, common.NewResult().Fail(401, "手机号格式不正确"))
	}
	//2.查看验证码是否存在
	if code, err := h.svcCtx.RedisCache.Get(getKey(PHONE_CODE_PREFIX, req.Phone)); err != nil || len(code) != 0 {
		httpx.OkJsonCtx(r.Context(), w, common.NewResult().Fail(405, "发送验证码错误"))
	}

	//3.todo：查看手机号状态是否封禁状态

	//4.调用rpc接口发送验证码
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	res, err := h.svcCtx.RegisterRpc.SendCode(ctx, &register.CodeReq{Phone: req.Phone})
	if err == nil {
		httpx.OkJson(w, common.NewResult().Success(res))
	}
	//4.错误为空则返回发送验证码失败
	httpx.OkJson(w, common.NewResult().Fail(406, err.Error()))
}

func getKey(prefix string, key string) string {
	return prefix + key
}
