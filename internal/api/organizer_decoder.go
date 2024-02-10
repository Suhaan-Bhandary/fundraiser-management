package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
)

func decodeRegisterOrganizerRequest(r *http.Request) (dto.RegisterOrganizerRequest, error) {
	var req dto.RegisterOrganizerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.RegisterOrganizerRequest{}, errors.New("Invalid Json in request body.")
	}

	return req, nil
}

func decodeLoginOrganizerRequest(r *http.Request) (dto.LoginOrganizerRequest, error) {
	var req dto.LoginOrganizerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.LoginOrganizerRequest{}, errors.New("Invalid Json in request body.")
	}

	return req, nil
}

func decodeUpdateOrganizerRequest(r *http.Request) (dto.UpdateOrganizerRequest, error) {
	var req dto.UpdateOrganizerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.UpdateOrganizerRequest{}, errors.New("Invalid Json in request body.")
	}

	tokenData, err := decodeTokenFromContext(r.Context())
	if err != nil {
		return dto.UpdateOrganizerRequest{}, internal_errors.InvalidCredentialError{Message: "Organizer not found"}
	}
	req.OrganizerId = uint(tokenData.ID)

	return req, nil
}
