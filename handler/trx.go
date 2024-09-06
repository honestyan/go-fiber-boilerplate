package handler

import (
	"github.com/gofiber/fiber/v2"
	U "github.com/honestyan/go-fiber-boilerplate/utils"
)

func rollbackCtxTrx(ctx *fiber.Ctx) {
	trx, _ := U.StartNewPGTrx(ctx)

	if trx != nil {
		if err := trx.Rollback(); err != nil {
			U.Log.Errorf("Error rolling back transaction: %v", err)
		}
	}
}

func commitCtxTrx(ctx *fiber.Ctx) error {
	trx, err := U.StartNewPGTrx(ctx)

	if err != nil {
		msg := "Unable to get transaction"
		return BuildError(ctx, msg, fiber.StatusInternalServerError, err)
	}

	if trx != nil {
		if err := trx.Commit(); err != nil {
			U.Log.Errorf("Error committing transaction: %v", err)
			return BuildError(ctx, "Error committing transaction", fiber.StatusInternalServerError, err)
		}
	}

	return nil
}
