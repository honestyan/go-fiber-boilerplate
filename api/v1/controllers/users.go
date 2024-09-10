package controllers

import (
	"github.com/gofiber/fiber/v2"
	"time"

	S "github.com/honestyan/go-fiber-boilerplate/api/v1/services"
	H "github.com/honestyan/go-fiber-boilerplate/handler"
	U "github.com/honestyan/go-fiber-boilerplate/utils"
	T "github.com/honestyan/go-fiber-boilerplate/api/v1/types" 
	"github.com/golang-jwt/jwt/v5"
)

func GetUsers(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	users, serviceErr := S.GetUsers(dbTrx, ctx.UserContext())

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"success": true,
		"users": users,
	})
}

func GetUser(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	idInt, err := ctx.ParamsInt("id")

	if err != nil {
		return H.BuildError(ctx, "Invalid user id", fiber.StatusBadRequest, err)
	}

	user, serviceErr := S.GetUser(dbTrx, ctx.UserContext(), idInt)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"success": true,
		"user": user,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	body := new(T.UserBody)

	if err := ctx.BodyParser(body); err != nil {
		return H.BuildError(ctx, "Invalid request body", fiber.StatusBadRequest, err)
	}

	user, serviceErr := S.CreateUser(dbTrx, ctx.UserContext(), body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	user.Password = ""

	return H.Success(ctx, fiber.Map{
		"success": true,
		"user": user,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	idInt, err := ctx.ParamsInt("id")

	if err != nil {
		return H.BuildError(ctx, "Invalid user id", fiber.StatusBadRequest, err)
	}

	body := new(T.UserBody)

	if err := ctx.BodyParser(body); err != nil {
		return H.BuildError(ctx, "Invalid request body", fiber.StatusBadRequest, err)
	}

	user, serviceErr := S.UpdateUser(dbTrx, ctx.UserContext(), idInt, body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	user.Password = ""

	return H.Success(ctx, fiber.Map{
		"success": true,
		"user": user,
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	idInt, err := ctx.ParamsInt("id")

	if err != nil {
		return H.BuildError(ctx, "Invalid user id", fiber.StatusBadRequest, err)
	}

	serviceErr := S.DeleteUser(dbTrx, ctx.UserContext(), idInt)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"success": true,
	})
}

func Login(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	body := new(T.LoginBody)

	if err := ctx.BodyParser(body); err != nil {
		return H.BuildError(ctx, "Invalid request body", fiber.StatusBadRequest, err)
	}

	user, serviceErr := S.Login(dbTrx, ctx.UserContext(), body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email": user.Email,
		"name": user.Name,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte("secret"))

	if err != nil {
		return H.BuildError(ctx, "Error generating token", fiber.StatusInternalServerError, err)
	}

	return H.Success(ctx, fiber.Map{
		"success": true,
		// "user": user,
		"token": token,
	})
}

