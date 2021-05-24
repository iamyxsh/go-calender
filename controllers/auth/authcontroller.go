package authcontroller

import (
	"calendly/database/Users"
	authreqbodies "calendly/schema/requestbodies/auth"
	"calendly/schema/resbodies"
	"calendly/utils/bcrypt"
	errorhandler "calendly/utils/error"

	"github.com/gofiber/fiber/v2"
)

func SignUpHandler(c *fiber.Ctx) error {
	user := new(authreqbodies.SignUpReq)

	errorhandler.CheckBodyParser(c, user)

	hashedPass, err := bcrypt.HashPass(user.Password)
	errorhandler.CheckErr(err, c)

	err = Users.CreateUser(user.Name, user.Email, hashedPass)
	errorhandler.CheckErr(err, c)

	return c.Status(fiber.StatusCreated).JSON(resbodies.SuccessRes("msg", "The user was created successfully."))
}
