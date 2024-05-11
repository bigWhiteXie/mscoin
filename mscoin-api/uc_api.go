package main

import (
	"flag"
	"fmt"
	"ucenter-api/internal/config"
	"ucenter-api/internal/handler"
	"ucenter-api/internal/router"
	"ucenter-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	routers := router.NewRouters(server)
	handler.RegisterMarketHandlers(routers, ctx)
	handler.RegisterUcenterHandlers(routers, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
