package dao

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"job-center/internal/model"
	"log"
)

type KlineDao struct {
	db *mongo.Database
}

func NewKlineDao(db *mongo.Database) *KlineDao {
	return &KlineDao{db: db}
}

func (k *KlineDao) DeleteGtTime(ctx context.Context, time int64, symbol string, period string) error {
	collection := k.db.Collection("exchange_kline_" + symbol + "_" + period)
	deleteResult, err := collection.DeleteMany(ctx, bson.D{{"time", bson.D{{"$gte", time}}}})
	log.Printf("删除表%s，数量为：%d \n", "exchange_kline_"+symbol+"_"+period, deleteResult.DeletedCount)
	return err
}

func (k *KlineDao) SaveBatch(ctx context.Context, klines []*model.Kline, symbol string, period string) error {
	if len(klines) == 0 {
		log.Printf("klines is empty")
		return errors.New("klines is empty")
	}
	collection := k.db.Collection("exchange_kline_" + symbol + "_" + period)

	//for i, k := range klines {
	//	if k == nil {
	//		log.Fatalf("Invalid Kline at index %d: nil", i) // 确保没有 nil
	//	}
	//}
	// 执行 InsertMany
	_, err := collection.InsertMany(context.Background(), toInterfaces(klines)) // 转换为 interface{}
	if err != nil {
		log.Fatalf("Failed to insert many: %v", err)
	}
	return err
}

// 转换结构体数组为接口数组
func toInterfaces(klines []*model.Kline) []interface{} {
	var ds []interface{}
	for _, k := range klines {
		if k != nil {
			ds = append(ds, k)
		}
	}
	return ds
}
