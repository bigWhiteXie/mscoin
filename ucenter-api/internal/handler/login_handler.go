package handler

import (
	"coin-common/common"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type LoginHandler struct {
	svcCtx *svc.ServiceContext
}

const ()

func NewLoginHandler(ctx *svc.ServiceContext) *LoginHandler {
	return &LoginHandler{svcCtx: ctx}
}

// 封装Login接口需要的中间件以及对应的路由，如果同一个业务有多个接口可以
func LoginHandlers(router *Routers, ctx *svc.ServiceContext) {
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
