package errorhandler

import (
	"calendly/schema/resbodies"
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func CheckErr(err error, c *fiber.Ctx) {
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(resbodies.FailRes("message", err.Error()))
		return
	}
}

func HandleErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ValidateStruct(body interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = fmt.Sprintf("%v", strings.Split(err.StructNamespace(), ".")[1])
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func CheckBodyParser(c *fiber.Ctx, schema interface{}) {
	err := c.BodyParser(schema)
	CheckErr(err, c)
}
