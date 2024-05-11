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

type MarketHandler struct {
	svcCtx *svc.ServiceContext
}

func (m *MarketHandler) GetCoinsThumbTrend(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq = types.MarketReq{}
	result := common.NewResult()

	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip

	l := logic.NewMarketLogic(r.Context(), m.svcCtx)
	resp, err := l.GetCoinsThumbTrend(&req)

	httpx.OkJsonCtx(r.Context(), w, result.Deal(resp, err))
}

func NewMarketHandler(svcCtx *svc.ServiceContext) *MarketHandler {
	return &MarketHandler{svcCtx}
}
