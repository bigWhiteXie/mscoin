package market

import (
	"ucenter-api/internal/router"
	"ucenter-api/internal/svc"
)

// 封装register接口需要的中间件以及对应的路由，如果同一个业务有多个接口可以
func RegisterHandlers(router *router.Routers, ctx *svc.ServiceContext) {
	handler := NewExchangeRateHandler(ctx)
	group := router.Group()

	group.Get("/market/exchange-rate/usd/:unit", handler.GetUsdRate)

}
