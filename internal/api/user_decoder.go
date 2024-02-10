package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

func decodeRegisterUserRequest(r *http.Request) (dto.RegisterUserRequest, error) {
	var req dto.RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.RegisterUserRequest{}, errors.New("invalid Json in request body")
	}

	return req, nil
}

func decodeLoginUserRequest(r *http.Request) (dto.LoginUserRequest, error) {
	var req dto.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.LoginUserRequest{}, errors.New("invalid Json in request body")
	}

	return req, nil
}
