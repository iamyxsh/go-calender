package rootrouter

import (
	authrouter "calendly/routes/auth"
	meetingrouter "calendly/routes/meeting"

	"github.com/gofiber/fiber/v2"
)

func RootRouter(api *fiber.Router) {
	a := *api

	auth := a.Group("auth")
	authrouter.AuthRouter(&auth)

	meeting := a.Group("meeting")
	meetingrouter.MettingRouter(&meeting)
}
