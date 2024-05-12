package main

import (
	"flag"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/zeromicro/go-zero/core/conf"
	"job-center/internal/config"
	"job-center/internal/market"
	"job-center/internal/svc"
	"time"
)

var configFile = flag.String("f", "./etc/conf.yaml", "the config file")

func main() {
	flag.Parse()

	c := &config.Config{}
	fmt.Print(*configFile + "\n")
	conf.MustLoad(*configFile, c)
	svr := svc.NewServiceContext(c)
	k := market.NewKline(svr)

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Do(func() {
		k.Do("1m")
	})
	s.Every(1).Hours().Do(func() {
		k.Do("1h")
	})
	s.StartBlocking()

	//k := market.NewKline(c)
	//k.Do("1m")
}
