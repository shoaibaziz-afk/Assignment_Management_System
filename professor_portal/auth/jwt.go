package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
GenerateToken creates a JWT token for professor authentication.
*/
func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  "professor",
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("dev-secret"))
}
