package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
)

func decodeRegisterOrganizerRequest(r *http.Request) (dto.RegisterOrganizerRequest, error) {
	var req dto.RegisterOrganizerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.RegisterOrganizerRequest{}, errors.New("invalid Json in request body")
	}

	return req, nil
}

func decodeLoginOrganizerRequest(r *http.Request) (dto.LoginOrganizerRequest, error) {
	var req dto.LoginOrganizerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.LoginOrganizerRequest{}, errors.New("invalid Json in request body")
	}

	return req, nil
}

func decodeUpdateOrganizerRequest(r *http.Request) (dto.UpdateOrganizerRequest, error) {
	var req dto.UpdateOrganizerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.UpdateOrganizerRequest{}, errors.New("invalid Json in request body")
	}

	tokenData, err := decodeTokenFromContext(r.Context())
	if err != nil {
		return dto.UpdateOrganizerRequest{}, internal_errors.InvalidCredentialError{Message: "Organizer not found"}
	}
	req.OrganizerId = uint(tokenData.ID)

	return req, nil
}

func decodeListOrganizerRequest(r *http.Request) (dto.ListOrganizersRequest, error) {
	value := r.URL.Query()

	offset, err := strconv.Atoi(value.Get("offset"))
	if err != nil || offset < 0 {
		return dto.ListOrganizersRequest{}, internal_errors.BadRequest{Message: "Invalid offset value"}
	}

	limit, err := strconv.Atoi(value.Get("limit"))
	if err != nil || limit < 0 {
		return dto.ListOrganizersRequest{}, internal_errors.BadRequest{Message: "Invalid limit value"}
	}

	// Keeping default as ascending order
	isAscending, err := strconv.ParseBool(value.Get("is_ascending"))
	if err != nil {
		isAscending = true
	}

	req := dto.ListOrganizersRequest{
		Search:             value.Get("search"),
		Verified:           value.Get("verified"),
		Offset:             uint(offset),
		Limit:              uint(limit),
		OrderByKey:         value.Get("order_by"),
		OrderByIsAscending: isAscending,
	}
	return req, nil
}
