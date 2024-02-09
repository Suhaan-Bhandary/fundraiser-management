package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
)

func decodeCreateFundraiser(r *http.Request) (dto.CreateFundraiserRequest, error) {
	var req dto.CreateFundraiserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.CreateFundraiserRequest{}, errors.New("Invalid Json in request body.")
	}

	return req, nil
}

func decodeCreateDonation(r *http.Request) (dto.CreateDonationRequest, error) {
	var req dto.CreateDonationRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.CreateDonationRequest{}, internal_errors.BadRequest{Message: "Invalid body data"}
	}

	tokenData, err := decodeTokenFromContext(r.Context())
	if err != nil {
		return dto.CreateDonationRequest{}, internal_errors.InvalidCredentialError{Message: "user id not found"}
	}

	fundraiserId, err := decodeId(r)
	if err != nil || fundraiserId <= 0 {
		return dto.CreateDonationRequest{}, internal_errors.BadRequest{Message: "Invalid fundraiser Id in URL"}
	}

	req.UserId = uint(tokenData.ID)
	req.FundraiserId = uint(fundraiserId)

	return req, nil
}
