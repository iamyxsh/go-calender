package middlewares

import (
	"calendly/database/Users"
	"calendly/schema/entities"
	"calendly/schema/resbodies"
	errorhandler "calendly/utils/error"
	"calendly/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUniqEmail(c *fiber.Ctx) error {
	user := new(entities.User)
	errorhandler.CheckBodyParser(c, user)

	_, err := Users.FindUserByAttr(bson.M{"email": user.Email})

	if err == nil {
		return c.Status(fiber.StatusForbidden).JSON(resbodies.FailRes("msg", "Email already exists."))
	}

	return c.Next()
}

func IsAuth(c *fiber.Ctx) error {
	userid, err := jwt.CheckToken(c)

	errorhandler.CheckErr(err, c)

	c.Locals("userid", userid)
	return c.Next()
}
