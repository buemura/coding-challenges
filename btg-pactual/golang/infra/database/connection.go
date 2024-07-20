package database

import (
	"context"
	"log"

	"github.com/buemura/btg-challenge/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client    *mongo.Client
	Ctx       = context.TODO()
	OrderColl *mongo.Collection
)

func Connect() {
	clientOptions := options.Client().ApplyURI(config.DATABASE_URL)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	OrderColl = client.Database("bgpactual").Collection("orders")
}
