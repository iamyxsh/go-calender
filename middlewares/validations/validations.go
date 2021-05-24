package validations

import (
	authreqbodies "calendly/schema/requestbodies/auth"
	errorhandler "calendly/utils/error"

	"github.com/gofiber/fiber/v2"
)

func CheckSignUpReq(c *fiber.Ctx) error {
	user := new(authreqbodies.SignUpReq)

	errorhandler.CheckBodyParser(c, user)

	errors := errorhandler.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
