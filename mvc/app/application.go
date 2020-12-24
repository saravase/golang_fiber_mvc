package app

import (
	fiber "github.com/gofiber/fiber/v2"
)

var (
	app = fiber.New()
)

func StartApplication() {

	mapURLs()

	app.Listen(":9090")
}
