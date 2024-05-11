package handler

import (
	"ucenter-api/internal/router"
	"ucenter-api/internal/svc"
)

// 封装register接口需要的中间件以及对应的路由，如果同一个业务有多个接口可以
func RegisterMarketHandlers(router *router.Routers, ctx *svc.ServiceContext) {
	rateHandler := NewExchangeRateHandler(ctx)
	marketHandler := NewMarketHandler(ctx)
	group := router.Group()
	group.Get("/market/exchange-rate/usd/:unit", rateHandler.GetUsdRate)
	group.Get("/market/symbol-thumb-trend", marketHandler.GetCoinsThumbTrend)
}
