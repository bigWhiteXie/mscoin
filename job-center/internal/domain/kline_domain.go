package domain

import (
	"context"
	"errors"
	"job-center/internal/dao"
	"job-center/internal/model"
)

type KlineDomain struct {
	klineDao *dao.KlineDao
}

func NewKlineDomain(klineDao *dao.KlineDao) *KlineDomain {
	return &KlineDomain{klineDao: klineDao}
}

func (d *KlineDomain) Save(data [][]string, symbol string, period string) error {
	if len(data) == 0 {
		return errors.New("data is empty")
	}
	//1.将data封装成kline
	klines := make([]*model.Kline, 20)
	for _, line := range data {
		kline := model.NewKline(line, period)
		klines = append(klines, kline)
	}
	ctx := context.Background()
	if err := d.klineDao.DeleteGtTime(ctx, klines[len(klines)-1].Time, symbol, period); err != nil {
		return err
	}

	return d.klineDao.SaveBatch(ctx, klines, symbol, period)
}
