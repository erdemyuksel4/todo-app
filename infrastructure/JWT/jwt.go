package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("secret_key")

func GenerateToken(userID int, userType int) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"type":   userType,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
