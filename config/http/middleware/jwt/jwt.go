package jwt

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JwtMiddleware() func(ctx *fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	}
	return jwtware.New(config)
}
