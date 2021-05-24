package resbodies

import "github.com/gofiber/fiber/v2"

func SuccessRes(key string, value interface{}) fiber.Map {
	return fiber.Map{
		"status": true,
		key:      value,
	}
}

func FailRes(key string, value interface{}) fiber.Map {
	return fiber.Map{
		"status": false,
		key:      value,
	}
}
