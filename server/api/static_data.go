package api

import (
	"context"
	"github.com/Sahil624/learn-docker/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func RegisterStaticRoute(app *fiber.App) {
	staticRoute := app.Group("/static_data")
	staticRoute.Get("/search", coinSearch)
}

type searchRequest struct {
	Query string `json:"query"`
}

func coinSearch(ctx *fiber.Ctx) error {
	var req searchRequest
	if err := ctx.QueryParser(&req); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"reason": "Invalid search Query",
		})
	}

	query := bson.D{}

	if req.Query != "" {
		query = bson.D{
			{"symbol_id", req.Query},
		}
	}

	db := database.GetDatabase()
	opt := options.Find()
	opt.SetLimit(30)
	response, err := db.Collection("coinDB").Find(context.Background(), query, opt)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"reason": "Internal error",
		})
	}

	var responseArray = []Symbol{}
	err = response.All(context.Background(), &responseArray)

	if err != nil {
		log.Println("Scanning error. -> ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"reason": "Internal error",
		})
	}

	return ctx.JSON(responseArray)
}
