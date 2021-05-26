package authcontroller

import (
	"calendly/database/Users"
	authreqbodies "calendly/schema/requestbodies/auth"
	"calendly/schema/resbodies"
	"calendly/utils/bcrypt"
	errorhandler "calendly/utils/error"
	"calendly/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func SignUpHandler(c *fiber.Ctx) error {
	user := c.Locals("body").(*authreqbodies.SignUpReq)

	hashedPass, err := bcrypt.HashPass(user.Password)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	err = Users.CreateUser(user.Name, user.Email, hashedPass)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	return c.Status(fiber.StatusCreated).JSON(resbodies.SuccessRes("msg", "The user was created successfully."))
}

func SignInHandler(c *fiber.Ctx) error {
	login := c.Locals("body").(*authreqbodies.SignInReq)

	user, err := Users.FindUserByAttr(bson.M{"email": login.Email})
	if err != nil {
		return c.Status(fiber.StatusAccepted).JSON(resbodies.FailRes("msg", "Account with this email not found."))
	}

	err = bcrypt.CompPass(login.Password, user.Password)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(resbodies.FailRes("msg", "Password does not match."))
	}

	token, e := jwt.CreateToken(user.ID.Hex())
	if e != nil {
		return errorhandler.SendErr(err, c)
	}
	user.Password = ""

	return c.Status(fiber.StatusAccepted).JSON(resbodies.SuccessRes("payload", fiber.Map{"token": token, "user": user}))
}
