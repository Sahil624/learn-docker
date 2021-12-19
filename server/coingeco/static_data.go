package coingeco

import (
	"context"
	"github.com/Sahil624/learn-docker/database"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	coingeckoType "github.com/superoo7/go-gecko/v3/types"
)

func getListOfCoins() (*coingeckoType.CoinList, error) {

	db := database.GetDatabase()
	count, err := db.Collection("coinDB").CountDocuments(context.Background(), bson.D{})

	if err != nil {
		log.Println("Error in counting symbols. Err :-", err)
		return nil, err
	}

	if count > 0 {
		log.Printf("Not adding static data. Already %d coins exists", count)
		return nil, nil
	}

	list, err := cgClient.CoinsList()

	if err != nil {
		log.Println("Error in fetching coin list", err)
	} else {
		log.Println("Found coins in market")
	}

	var data []interface{}

	for _, coin := range *list {
		data = append(data, bson.D{
			{"symbol_id", coin.ID},
			{"symbol", coin.Symbol},
			{"name", coin.Name},
		})
	}

	_, err = db.Collection("coinDB").InsertMany(
		context.Background(),
		data,
	)

	if err != nil {
		log.Println("Could not save data in DB", err)
		return list, err
	}

	return list, err
}
