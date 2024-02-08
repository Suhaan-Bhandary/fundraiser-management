package helpers

import (
	"time"

	"os"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func CreateToken(user_id int, role string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 30) // 30 Days expiration

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user_id,
			"role":    role,
			"exp":     expirationTime,
		})

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if err != nil || !token.Valid {
		return internal_errors.InvalidCredentialError{Message: "Invalid JWT token"}
	}

	return nil
}
