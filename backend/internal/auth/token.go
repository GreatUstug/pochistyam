package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret []byte

func InitSecret() error {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return errors.New("Environment variable JWT_SECRET is not set")
	}
	jwtSecret = []byte(secret)
	return nil
}

func GenerateToken(userID string) (string, error) {
	claims := &jwt.MapClaims{
		"user_id": userID,
		"jti":     uuid.NewString(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
