package domain

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"market/internal/config"
	"market/internal/model"
	"market/internal/svc"
	"testing"
)

func TestMarketDomain_SymbolThumbTrend(t *testing.T) {
	var configFile = flag.String("f", "../../etc/conf.yaml", "the config file")
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	domain := NewMarketDomain(ctx)
	coin := &model.ExchangeCoin{
		Symbol: "BTC/USDT",
	}
	trend, _ := domain.SymbolThumbTrend(coin)
	fmt.Print(trend)
}
