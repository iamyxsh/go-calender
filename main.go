package main

import (
	"calendly/constants/variables"
	"calendly/database"
	rootrouter "calendly/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	file, err := os.OpenFile("./logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	api := app.Group("api")
	rootrouter.RootRouter(&api)

	database.ConnectDB()
	app.Listen(variables.PORT)
}
