package middleware

import (
    "time"
    "github.com/gofiber/fiber/v2"
    U "github.com/honestyan/go-fiber-boilerplate/utils"
)

func LoggerMiddleware(c *fiber.Ctx) error {
    start := time.Now()
    
    err := c.Next()

    U.Log.Infof("Request: Method=%s, Path=%s, Status=%d, Duration=%v",
        c.Method(), c.Path(), c.Response().StatusCode(), time.Since(start))
    
    return err
}
