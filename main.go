package main

import (
	"calendly/constants/variables"
	"calendly/database"
	rootrouter "calendly/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("api")
	rootrouter.RootRouter(&api)

	database.ConnectDB()
	app.Listen(variables.PORT)
}
