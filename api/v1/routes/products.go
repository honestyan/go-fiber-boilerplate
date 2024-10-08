package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/honestyan/go-fiber-boilerplate/api/v1/controllers"
	mw "github.com/honestyan/go-fiber-boilerplate/api/v1/middleware"
	C "github.com/honestyan/go-fiber-boilerplate/constants"
	T "github.com/honestyan/go-fiber-boilerplate/api/v1/types"
)

func SetupProductsRoutes(router fiber.Router) {
	router.Get("/products", mw.RateLimit(C.Tier3, 0), mw.AuthMiddleware, controllers.GetProducts)
	
	router.Get("/products/:id", mw.RateLimit(C.Tier3, 0), mw.AuthMiddleware, controllers.GetProduct)

	router.Post("/products", mw.RateLimit(C.Tier2, 0), mw.AuthMiddleware, mw.ValidateRequestBody(T.ProductBody{}), mw.AuthMiddleware, controllers.CreateProduct)
	
	router.Patch("/products/:id", mw.RateLimit(C.Tier2, 0), mw.ValidateRequestBody(T.ProductBody{}), mw.AuthMiddleware, controllers.UpdateProduct)

	router.Delete("/products/:id", mw.RateLimit(C.Tier3, 0), mw.AuthMiddleware, controllers.DeleteProduct)

}
