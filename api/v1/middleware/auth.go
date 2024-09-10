package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	H "github.com/honestyan/go-fiber-boilerplate/handler"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return H.BuildError(c, "Unauthorized", fiber.StatusUnauthorized, nil)
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return H.BuildError(c, "Invalid token", fiber.StatusUnauthorized, err)
	}

	if !token.Valid {
		return H.BuildError(c, "Token is not valid", fiber.StatusUnauthorized, nil)
	}

	c.Locals("user", claims)

	return c.Next()
}

