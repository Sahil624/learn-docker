package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var client *mongo.Client

func ConnectMongo() error {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://admin:password@localhost:27017/"))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	return err
}

func GetDatabase() *mongo.Database {
	return client.Database("cryptoDB")
}
