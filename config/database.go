package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func New() *mongo.Database {
	if db != nil {
		return db
	}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mbeledos:mbeledos123@ds043991.mlab.com:43991/mbeledos").SetRetryWrites(false))
	if err != nil {
		log.Fatal("Can't Connect to MongoDB", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Timeout more than 10 sec", err.Error())
	}

	db = client.Database("mbeledos")
	return db
}
