package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
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
