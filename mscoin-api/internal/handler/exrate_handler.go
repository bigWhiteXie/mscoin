package handler

import (
	"coin-common/common"
	"coin-common/tools"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	types "ucenter-api/internal/types/market"
)

type ExchangeRateHandler struct {
	svcCtx *svc.ServiceContext
}

func (h *ExchangeRateHandler) GetUsdRate(w http.ResponseWriter, r *http.Request) {
	var req types.RateRequest
	if err := httpx.ParsePath(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewExchangeRateLogic(r.Context(), h.svcCtx)
	resp, err := l.GetUsdRate(&req)
	result := common.NewResult().Deal(resp.Rate, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func NewExchangeRateHandler(svcCtx *svc.ServiceContext) *ExchangeRateHandler {
	return &ExchangeRateHandler{svcCtx}
}
