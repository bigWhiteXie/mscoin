package dao

import (
	"context"
	"gorm.io/gorm"
	"market/internal/model"
)

type ExchangeCoinDao struct {
	db *gorm.DB
}

func NewExchangeCoinDao(db *gorm.DB) *ExchangeCoinDao {
	return &ExchangeCoinDao{db: db}
}

func (d *ExchangeCoinDao) FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error) {
	err = d.db.Model(&model.ExchangeCoin{}).Where("visible=?", 1).Find(&list).Error
	return
}
