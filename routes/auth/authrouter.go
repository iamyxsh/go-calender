package authrouter

import (
	authcontroller "calendly/controllers/auth"
	"calendly/middlewares"
	"calendly/middlewares/validations"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(auth *fiber.Router) {
	a := *auth

	a.Post("/signup", validations.CheckSignUpReq, middlewares.CheckUniqEmail, authcontroller.SignUpHandler)
}
