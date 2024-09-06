package controllers

import (
	C "github.com/honestyan/go-fiber-boilerplate/config"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"v":   C.Conf.Version,
		"env": C.Conf.Environment,
	})
}
