package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"market/internal/config"
	"time"
)

type MongoClient struct {
	cli *mongo.Client
	Db  *mongo.Database
}

func ConnectMongo(c *config.Config) *MongoClient {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(c.Mongo.Url))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	database := client.Database("mscoin")
	return &MongoClient{cli: client, Db: database}
}

func (c *MongoClient) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := c.cli.Disconnect(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println("关闭连接..")
}
