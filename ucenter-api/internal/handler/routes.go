// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

type Routers struct {
	server      *rest.Server
	middlewares []rest.Middleware
}

func NewRouters(svr *rest.Server) *Routers {
	return &Routers{server: svr}
}

func (r *Routers) Get(path string, h http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(r.middlewares,
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: h,
			},
		),
	)
}

func (r *Routers) Post(path string, h http.HandlerFunc) {
	r.server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    path,
			Handler: h,
		})
}

// 创建一个新的router，使得midware可以针对不同路径
func (r *Routers) Group() *Routers {
	return &Routers{server: r.server}
}

func (r *Routers) Use(middlewares ...rest.Middleware) {
	r.middlewares = middlewares
}
