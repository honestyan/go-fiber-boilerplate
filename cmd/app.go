package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/honestyan/go-fiber-boilerplate/api/v1/routes"
	H "github.com/honestyan/go-fiber-boilerplate/handler"
)

func InitApp() *fiber.App {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: H.ErrorHandler,
		},
	)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH, HEAD",
	}))

	app.Use(requestid.New())

	routes.SetupRoutes(app)

	return app
}
