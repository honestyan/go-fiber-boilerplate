package routes

import (
	"github.com/honestyan/go-fiber-boilerplate/api/v1/controllers"
	"github.com/gofiber/fiber/v2"
	mw "github.com/honestyan/go-fiber-boilerplate/api/v1/middleware"
)

func SetupRoutes(app *fiber.App) {
	
	app.Get("/", controllers.Health)
	
	v1API := app.Group("/api/v1")
	
	app.Use(mw.LoggerMiddleware)
	SetupProductsRoutes(v1API)
	SetupUsersRoutes(v1API)
}
