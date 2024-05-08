package domain

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"market/internal/dao"
	"market/internal/model"
)

type ExchangeCoinDomain struct {
	ExchangeCoinDao *dao.ExchangeCoinDao
}

func NewExchangeCoinDomain(db *gorm.DB) *ExchangeCoinDomain {
	return &ExchangeCoinDomain{
		ExchangeCoinDao: dao.NewExchangeCoinDao(db),
	}
}

func (d *ExchangeCoinDomain) FindVisible(ctx context.Context) []*model.ExchangeCoin {
	list, err := d.ExchangeCoinDao.FindVisible(ctx)
	if err != nil {
		logx.Error(err)
		return []*model.ExchangeCoin{}
	}
	return list
}
