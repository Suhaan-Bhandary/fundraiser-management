package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

func decodeLoginAdminRequest(r *http.Request) (dto.LoginAdminRequest, error) {
	var req dto.LoginAdminRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.LoginAdminRequest{}, errors.New("invalid Json in request body")
	}

	return req, nil
}
