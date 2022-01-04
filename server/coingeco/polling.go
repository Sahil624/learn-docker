package coingeco

import (
	"context"
	"fmt"
	"github.com/Sahil624/learn-docker/database"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

var polledSymbols = map[string]int32{}
var vc = []string{"usd"}

func startPolling() {
	ticker := time.Tick(time.Second * 30)

	for range ticker {
		fmt.Println("Tick", polledSymbols)
		fetchPrice()
	}
}

func fetchPrice() {
	db := database.GetDatabase()

	keys := make([]string, 0, len(polledSymbols))
	for k := range polledSymbols {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		return
	}

	response, err := cgClient.SimplePrice(keys, vc)

	if err != nil {
		log.Println("Error in getting symbol prices. Error :-", err)
	}
	var priceMap = map[string]float32{}
	for _, symbol := range keys {
		if symbolPrice, ok := (*response)[symbol]; ok {
			price := symbolPrice["usd"]
			priceMap[symbol] = symbolPrice["usd"]
			fmt.Printf("Symbol %s - Price %f\n", symbol, price)

			_, err = db.Collection("coinDB").UpdateOne(
				context.Background(),
				bson.D{{"symbol_id", symbol}},
				bson.D{{"$set", bson.D{{"ltp", price}}}},
			)
			if err != nil {
				log.Println("caught exception during transaction, aborting.")
			}
		}
	}
	GetChannel() <- priceMap
}

func AddToPoll(id string) {
	if val, ok := polledSymbols[id]; ok {
		polledSymbols[id] = val + 1
	} else {
		polledSymbols[id] = 1
	}
	fetchPrice()
}

func RemoveFromPoll(id string) {
	if val, ok := polledSymbols[id]; ok {
		polledSymbols[id] = val - 1
		if polledSymbols[id] == 0 {
			delete(polledSymbols, id)
		}
	}
}
