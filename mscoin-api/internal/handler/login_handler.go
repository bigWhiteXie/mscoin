package handler

import (
	"coin-common/common"
	"coin-common/tools"
	"github.com/zeromicro/go-zero/rest/httpx"
	"k8s.io/kube-openapi/pkg/validation/errors"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/router"
	"ucenter-api/internal/svc"
	types "ucenter-api/internal/types/ucenter"
)

type LoginHandler struct {
	svcCtx *svc.ServiceContext
}

const ()

func NewLoginHandler(ctx *svc.ServiceContext) *LoginHandler {
	return &LoginHandler{svcCtx: ctx}
}

// 封装Login接口需要的中间件以及对应的路由，如果同一个业务有多个接口可以
func LoginHandlers(router *router.Routers, ctx *svc.ServiceContext) {
	handler := NewLoginHandler(ctx)
	group := router.Group()
	group.Post("/uc/Login", handler.Login)
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req *types.LoginReq
	result := &common.Result{}
	if err := httpx.Parse(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	loginLogic := logic.NewLoginLogic(r.Context(), h.svcCtx)

	resp, err := loginLogic.Login(req)
	if err != nil {
		httpx.OkJson(w, result.Fail(404, err.Error()))
		return
	}

	httpx.OkJson(w, result.Success(resp))
}

func (h *LoginHandler) CheckLogin(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-auth-token")
	secret := h.svcCtx.Config.Jwt.AccessSecret
	result := &common.Result{}
	if userId, err := tools.ParseToken(token, secret); err == nil {
		httpx.OkJson(w, result.Success(userId))
		return
	}
	httpx.ErrorCtx(r.Context(), w, errors.New(401, "token无效"))
}
