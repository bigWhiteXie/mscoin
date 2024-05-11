package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"market/internal/model"
	"strconv"
)

type KlineDao struct {
	db *mongo.Database
}

func NewKlineDao(db *mongo.Database) *KlineDao {
	return &KlineDao{db: db}
}

func (k *KlineDao) FindKlineInTime(symbol string, period string, from int64, to int64) (list []*model.Kline, err error) {
	ctx := context.Background()
	name := fmt.Sprintf("exchange_kline_%s_%s", symbol, period)
	fmt.Println(name)
	fmt.Println(from)
	fmt.Println(to)

	collection := k.db.Collection(name)
	cur, err := collection.Find(ctx, bson.D{{"time", bson.D{{"$gte", from}, {"$lte", to}}}},
		&options.FindOptions{
			Sort: bson.D{{"time", -1}},
		})
	fmt.Println("length: " + strconv.Itoa(cur.RemainingBatchLength()))

	cur.All(ctx, &list)
	return
}
