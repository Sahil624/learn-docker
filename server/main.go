package main

import (
	"github.com/Sahil624/learn-docker/api"
	"github.com/Sahil624/learn-docker/coingeco"
	"github.com/Sahil624/learn-docker/database"
	"github.com/Sahil624/learn-docker/ticker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"log"
)

func main() {
	log.Println("Starting Server!!!")

	if err := database.ConnectMongo(); err != nil {
		log.Fatalln("Error in connecting database", err)
		return
	} else {
		log.Println("DB connected")
	}

	coingeco.UpdateCoins()

	app := fiber.New()

	app.Static("/", "./public_build")

	app.Use(cors.New())

	apiRoute := app.Group("/api")

	apiRoute.Use("/ws", ticker.UpgradeSocket)
	apiRoute.Use("/ws/ticker", websocket.New(ticker.WebsocketHandler))

	api.RegisterStaticRoute(apiRoute)
	address := ":3000"
	log.Println("Go Server Running on Address", address)
	app.Listen(address)
}
