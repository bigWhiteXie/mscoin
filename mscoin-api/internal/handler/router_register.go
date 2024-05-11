package handler

import (
	"ucenter-api/internal/router"
	"ucenter-api/internal/svc"
)

// 封装register接口需要的中间件以及对应的路由，如果同一个业务有多个接口可以
func RegisterUcenterHandlers(router *router.Routers, ctx *svc.ServiceContext) {
	registerHandler := NewRegisterHandler(ctx)
	loginHandler := NewLoginHandler(ctx)
	group := router.Group()

	group.Post("/uc/register", registerHandler.Register)
	group.Get("/uc/code", registerHandler.SendCode)
	group.Post("/uc/ucenter", loginHandler.Login)
}
