package validations

import (
	authreqbodies "calendly/schema/requestbodies/auth"
	meetingreqbodies "calendly/schema/requestbodies/meeting"
	"calendly/schema/resbodies"
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

	c.Locals("body", user)

	return c.Next()
}

func CheckSignInReq(c *fiber.Ctx) error {
	loginbody := new(authreqbodies.SignInReq)

	errorhandler.CheckBodyParser(c, loginbody)

	errors := errorhandler.ValidateStruct(*loginbody)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resbodies.FailRes("errors", fiber.Map{
			"errors": errors,
		}))
	}

	c.Locals("body", loginbody)

	return c.Next()
}

func CheckCreateMeetingReq(c *fiber.Ctx) error {
	createBody := new(meetingreqbodies.CreateMeetingReq)
	errorhandler.CheckBodyParser(c, createBody)

	errors := errorhandler.ValidateStruct(*createBody)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resbodies.FailRes("errors", fiber.Map{
			"errors": errors,
		}))
	}

	c.Locals("body", createBody)

	return c.Next()
}

func CheckCreateSlotReq(c *fiber.Ctx) error {
	createBody := new(meetingreqbodies.CreateSlotReq)
	errorhandler.CheckBodyParser(c, createBody)

	errors := errorhandler.ValidateStruct(*createBody)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resbodies.FailRes("errors", fiber.Map{
			"errors": errors,
		}))
	}

	c.Locals("body", createBody)

	return c.Next()
}
