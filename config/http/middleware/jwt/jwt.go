package jwt

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func JwtMiddleware() func(ctx *fiber.Ctx) error {
	godotenv.Load(".env")
	config := jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	}
	return jwtware.New(config)
}
