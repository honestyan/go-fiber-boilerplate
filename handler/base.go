package handler

import (
    "github.com/gofiber/fiber/v2"
    U "github.com/honestyan/go-fiber-boilerplate/utils"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
    U.Log.Errorf("Error occurred: %v", err)
    return BuildError(ctx, "Internal Server Error", fiber.StatusInternalServerError, err)
}

func BuildError(ctx *fiber.Ctx, message interface{}, code int, originalErr error) error {
    rollbackCtxTrx(ctx)
    if code == 0 {
        code = fiber.StatusInternalServerError
    }
    var detail string
    if originalErr != nil {
        detail = originalErr.Error()
        U.Log.Errorf("BuildError: %s - %v", message, originalErr)
    }
    return ctx.Status(code).JSON(fiber.Map{
        "success": false,
        "message": message,
        "detail":  detail,
    })
}

func Success(ctx *fiber.Ctx, data interface{}) error {
    err := commitCtxTrx(ctx)
    if err != nil {
        U.Log.Errorf("Commit transaction error: %v", err)
        return err
    }
    // U.Log.Info("Request successful")
    return ctx.JSON(data)
}
