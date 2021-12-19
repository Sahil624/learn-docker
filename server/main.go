package main

import (
	"github.com/Sahil624/learn-docker/api"
	"github.com/Sahil624/learn-docker/coingeco"
	"github.com/Sahil624/learn-docker/database"
	"github.com/Sahil624/learn-docker/ticker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func main() {
	log.Println("Starting Server")

	if err := database.ConnectMongo(); err != nil {
		log.Fatalln("Error in connecting database", err)
		return
	} else {
		log.Println("DB connected")
	}

	coingeco.UpdateCoins()

	app := fiber.New()

	app.Use("/ws", ticker.UpgradeSocket)
	app.Use("/ws/ticker", websocket.New(ticker.WebsocketHandler))

	api.RegisterStaticRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
