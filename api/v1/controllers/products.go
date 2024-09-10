package controllers

import (
	"github.com/gofiber/fiber/v2"

	S "github.com/honestyan/go-fiber-boilerplate/api/v1/services"
	H "github.com/honestyan/go-fiber-boilerplate/handler"
	U "github.com/honestyan/go-fiber-boilerplate/utils"
	T "github.com/honestyan/go-fiber-boilerplate/api/v1/types" 
)

func GetProducts(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	products, serviceErr := S.GetProducts(dbTrx, ctx.UserContext())

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"success": true,
		"data":    products,
	})
}

func GetProduct(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	idInt, err := ctx.ParamsInt("id")

	if err != nil {
		return H.BuildError(ctx, "Invalid product id", fiber.StatusBadRequest, err)
	}

	product, serviceErr := S.GetProduct(dbTrx, ctx.UserContext(), idInt)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"success": true,
		"data":    product,
	})
}

func CreateProduct(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	body := &T.ProductBody{}

	if err := ctx.BodyParser(body); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}

	product, serviceErr := S.CreateProduct(dbTrx, ctx.UserContext(), body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"success": true,
		"data":	product,
	})
}

func UpdateProduct(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	idInt, err := ctx.ParamsInt("id")

	if err != nil {
		return H.BuildError(ctx, "Invalid product id", fiber.StatusBadRequest, err)
	}

	body := &T.ProductBody{}

	if err := ctx.BodyParser(body); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}

	product, serviceErr := S.UpdateProduct(dbTrx, ctx.UserContext(), idInt, body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"success": true,
		"data":	product,
	})
}

func DeleteProduct(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	idInt, err := ctx.ParamsInt("id")

	if err != nil {
		return H.BuildError(ctx, "Invalid product id", fiber.StatusBadRequest, err)
	}

	serviceErr := S.DeleteProduct(dbTrx, ctx.UserContext(), idInt)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"success": true,
	})
}
