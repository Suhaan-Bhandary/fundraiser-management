package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/helpers"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth(handler func(w http.ResponseWriter, r *http.Request), allowed []string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			ErrorResponse(w, http.StatusUnauthorized, errors.New("Missing authorization header"))
			return
		}

		tokenString = tokenString[len("Bearer "):]
		token, err := helpers.VerifyToken(tokenString)
		if err != nil {
			fmt.Println(err)
			ErrorResponse(w, http.StatusUnauthorized, errors.New("Invalid token"))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// jwt is storing the id as float64 so first take it as float and then convert it
			id := int(claims["id"].(float64))
			tokenRole := claims["role"].(string)

			for _, role := range allowed {
				if role == tokenRole {
					rcopy := r.WithContext(context.WithValue(r.Context(), "token-data", dto.Token{
						ID:   id,
						Role: tokenRole,
					}))

					handler(w, rcopy)
					return
				}
			}
		}

		ErrorResponse(w, http.StatusUnauthorized, errors.New("Invalid token"))
	}
}
