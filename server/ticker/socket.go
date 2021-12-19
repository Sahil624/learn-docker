package ticker

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func UpgradeSocket(ctx *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(ctx) {
		ctx.Locals("allowed", true)
		return ctx.Next()
	}
	return fiber.ErrUpgradeRequired
}

func WebsocketHandler(conn *websocket.Conn) {
	var (
		msg []byte
		err error
	)

	for {
		if _, msg, err = conn.ReadMessage(); err != nil {
			log.Println("read:", err)
			disconnected(conn)
			conn.Close()
			break
		}
		readSocketMessage(conn, msg)
	}
}
