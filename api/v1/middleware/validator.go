package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateRequestBody(expectedType interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := reflect.New(reflect.TypeOf(expectedType)).Interface()
		if err := c.BodyParser(payload); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}

		if err := validate.Struct(payload); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		c.Locals("validatedPayload", payload)

		return c.Next()
	}
}