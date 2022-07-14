package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(username string) (string, error) {
	claim := jwt.MapClaims{}
	claim["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claim["username"] = username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
