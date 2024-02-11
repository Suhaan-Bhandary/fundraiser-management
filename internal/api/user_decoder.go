package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
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

func decodeListUserRequest(r *http.Request) (dto.ListUserRequest, error) {
	value := r.URL.Query()

	offset, err := strconv.Atoi(value.Get("offset"))
	if err != nil || offset < 0 {
		return dto.ListUserRequest{}, internal_errors.BadRequest{Message: "Invalid offset value"}
	}

	limit, err := strconv.Atoi(value.Get("limit"))
	if err != nil {
		return dto.ListUserRequest{}, internal_errors.BadRequest{Message: "Invalid limit value"}
	}

	// Keeping default as ascending order
	isAscending, err := strconv.ParseBool(value.Get("is_ascending"))
	if err != nil {
		isAscending = true
	}

	req := dto.ListUserRequest{
		Search:             value.Get("search"),
		Offset:             uint(offset),
		Limit:              uint(limit),
		OrderByKey:         value.Get("order_by"),
		OrderByIsAscending: isAscending,
	}
	return req, nil
}
