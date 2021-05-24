package rootrouter

import (
	authrouter "calendly/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func RootRouter(api *fiber.Router) {
	a := *api

	auth := a.Group("auth")
	authrouter.AuthRouter(&auth)
}
