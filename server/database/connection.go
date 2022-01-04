package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var client *mongo.Client

func ConnectMongo() error {
	var err error
	//host := os.Getenv("HOSTNAME")
	//if host == "" {
	//	host = "localhost"
	//}
	host := "mongo"
	address := fmt.Sprintf("mongodb://admin:password@%s:27017/", host)
	fmt.Println("Mongo Address", address)
	clientOptions := options.Client().ApplyURI(address)
	client, err = mongo.Connect(context.Background(), clientOptions)
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
