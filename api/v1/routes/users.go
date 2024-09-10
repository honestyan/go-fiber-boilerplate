package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/honestyan/go-fiber-boilerplate/api/v1/controllers"
	mw "github.com/honestyan/go-fiber-boilerplate/api/v1/middleware"
	C "github.com/honestyan/go-fiber-boilerplate/constants"
	T "github.com/honestyan/go-fiber-boilerplate/api/v1/types"
)

func SetupUsersRoutes(router fiber.Router) {
	router.Get("/users", mw.RateLimit(C.Tier3, 0), mw.AuthMiddleware, controllers.GetUsers)
	
	router.Get("/users/:id", mw.RateLimit(C.Tier3, 0), mw.AuthMiddleware, controllers.GetUser)

	router.Post("/users", 
		mw.RateLimit(C.Tier2, 0), 
		mw.ValidateRequestBody(T.UserBody{}),
		mw.AuthMiddleware,
		controllers.CreateUser,
	)

	router.Patch("/users/:id", 
		mw.RateLimit(C.Tier2, 0), 
		mw.ValidateRequestBody(T.UserBody{}),
		mw.AuthMiddleware,
		controllers.UpdateUser,
	)

	router.Delete("/users/:id", mw.RateLimit(C.Tier3, 0), mw.AuthMiddleware, controllers.DeleteUser)

	router.Post("/users/login", mw.RateLimit(C.Tier3, 0), mw.ValidateRequestBody(T.LoginBody{}), controllers.Login)
}