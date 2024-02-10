package helpers

import (
	"fmt"
	"time"

	"os"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func CreateToken(id uint, role string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 30) // 30 Days expiration

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":   id,
			"role": role,
			"exp":  expirationTime.Unix(),
		})

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if err != nil || !token.Valid {
		fmt.Println(err, token.Valid)
		return nil, internal_errors.InvalidCredentialError{Message: "Invalid JWT token"}
	}

	return token, nil
}
