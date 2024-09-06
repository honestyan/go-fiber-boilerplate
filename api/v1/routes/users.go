package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/honestyan/go-fiber-boilerplate/api/v1/controllers"
	mw "github.com/honestyan/go-fiber-boilerplate/api/v1/middleware"
	C "github.com/honestyan/go-fiber-boilerplate/constants"
	T "github.com/honestyan/go-fiber-boilerplate/api/v1/types"
)

func SetupUsersRoutes(router fiber.Router) {
	router.Get("/users", mw.RateLimit(C.Tier3, 0), controllers.GetUsers)
	
	router.Get("/users/:id", mw.RateLimit(C.Tier3, 0), controllers.GetUser)

	router.Post("/users", 
		mw.RateLimit(C.Tier2, 0), 
		mw.ValidateRequestBody(T.UserBody{}),
		controllers.CreateUser,
	)

	router.Patch("/users/:id", 
		mw.RateLimit(C.Tier2, 0), 
		mw.ValidateRequestBody(T.UserBody{}),
		controllers.UpdateUser,
	)

	router.Delete("/users/:id", mw.RateLimit(C.Tier3, 0), controllers.DeleteUser)
}