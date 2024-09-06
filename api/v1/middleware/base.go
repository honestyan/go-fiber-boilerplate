package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	H "github.com/honestyan/go-fiber-boilerplate/handler"
)

func RateLimit(count int, duration time.Duration) fiber.Handler {

	if duration == 0 {
		duration = time.Minute
	}
	return limiter.New(limiter.Config{
		Max:        count,
		Expiration: duration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() + "_" + c.Path()
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return H.BuildError(ctx, "Too many requests!", fiber.ErrTooManyRequests.Code, nil)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
	})
}
