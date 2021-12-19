package ticker

import (
	"context"
	"fmt"
	"github.com/Sahil624/learn-docker/coingeco"
	"github.com/Sahil624/learn-docker/database"
	"github.com/gofiber/websocket/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type findSymbolResult struct {
	SymbolID string `bson:"symbol_id"`
}

var connectionSymbol = map[*websocket.Conn][]string{}
var symbolConnection = map[string][]*websocket.Conn{}

func readSocketMessage(conn *websocket.Conn, message []byte) {
	symbolIds := strings.Split(string(message), ",")
	fmt.Println("Read message", symbolIds)

	db := database.GetDatabase()

	query := bson.D{
		{
			"symbol_id", bson.D{
				{"$in", symbolIds},
			},
		},
	}

	findOptions := options.Find()
	findOptions.SetProjection(bson.M{
		"symbol_id": 1,
		"_id":       0,
	})

	cur, err := db.Collection("coinDB").Find(
		context.Background(), query, findOptions,
	)

	if err != nil {
		fmt.Println("Error in filtering socket symbols. Error :- ", err)
		return
	}

	var res []findSymbolResult
	err = cur.All(context.Background(), &res)

	if err != nil {
		fmt.Println("Error in finding socket symbols. Error :- ", err)
		return
	}

	connectionSymbol[conn] = []string{}
	for _, symbol := range res {
		connectionSymbol[conn] = append(connectionSymbol[conn], symbol.SymbolID)
		coingeco.AddToPoll(symbol.SymbolID)
		if _, ok := symbolConnection[symbol.SymbolID]; !ok {
			symbolConnection[symbol.SymbolID] = []*websocket.Conn{}
		}
		symbolConnection[symbol.SymbolID] = append(symbolConnection[symbol.SymbolID], conn)
	}
}

func disconnected(conn *websocket.Conn) {
	if _, ok := connectionSymbol[conn]; !ok {
		return
	}

	for _, symbol := range connectionSymbol[conn] {
		coingeco.RemoveFromPoll(symbol)
	}
}
