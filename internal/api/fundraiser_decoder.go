package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
)

func decodeCreateFundraiser(r *http.Request) (dto.CreateFundraiserRequest, error) {
	var req dto.CreateFundraiserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.CreateFundraiserRequest{}, errors.New("invalid Json in request body")
	}

	// Assigning the organizer id before validating it
	tokenData, err := decodeTokenFromContext(r.Context())
	if err != nil {
		return dto.CreateFundraiserRequest{}, internal_errors.InvalidCredentialError{Message: "Token not found"}
	}
	req.OrganizerId = tokenData.ID

	return req, nil
}

func decodeDeleteFundraiser(r *http.Request) (dto.DeleteFundraiserRequest, error) {
	fundraiserId, err := decodeId(r)
	if err != nil {
		return dto.DeleteFundraiserRequest{}, internal_errors.BadRequest{Message: "Fundraiser id not found"}
	}

	token, err := decodeTokenFromContext(r.Context())
	if err != nil {
		return dto.DeleteFundraiserRequest{}, internal_errors.InvalidCredentialError{Message: "Token not found"}
	}

	return dto.DeleteFundraiserRequest{
		FundraiserId: fundraiserId,
		Token:        token,
	}, nil
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

func decodeUpdateFundraiser(r *http.Request) (dto.UpdateFundraiserRequest, error) {
	var req dto.UpdateFundraiserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return dto.UpdateFundraiserRequest{}, internal_errors.BadRequest{Message: "Invalid body data"}
	}

	tokenData, err := decodeTokenFromContext(r.Context())
	if err != nil {
		return dto.UpdateFundraiserRequest{}, internal_errors.InvalidCredentialError{Message: "organizer not found"}
	}

	fundraiserId, err := decodeId(r)
	if err != nil || fundraiserId <= 0 {
		return dto.UpdateFundraiserRequest{}, internal_errors.BadRequest{Message: "Invalid fundraiser Id in URL"}
	}

	req.FundraiserId = uint(fundraiserId)
	req.RequestOrganizerId = uint(tokenData.ID)

	return req, nil
}

func decodeUserDonationsRequest(r *http.Request) (dto.ListUserDonationsRequest, error) {
	value := r.URL.Query()

	offset, err := strconv.Atoi(value.Get("offset"))
	if err != nil || offset < 0 {
		return dto.ListUserDonationsRequest{}, internal_errors.BadRequest{Message: "Invalid offset value"}
	}

	limit, err := strconv.Atoi(value.Get("limit"))
	if err != nil || limit < 0 {
		return dto.ListUserDonationsRequest{}, internal_errors.BadRequest{Message: "Invalid limit value"}
	}

	// Keeping default as ascending order
	isAscending, err := strconv.ParseBool(value.Get("is_ascending"))
	if err != nil {
		isAscending = true
	}

	tokenData, err := decodeTokenFromContext(r.Context())
	if err != nil {
		return dto.ListUserDonationsRequest{}, internal_errors.InvalidCredentialError{Message: "Token not found"}
	}

	req := dto.ListUserDonationsRequest{
		UserId:             tokenData.ID,
		Search:             value.Get("search"),
		IsAnonymous:        value.Get("is_anonymous"),
		Offset:             uint(offset),
		Limit:              uint(limit),
		OrderByKey:         value.Get("order_by"),
		OrderByIsAscending: isAscending,
	}
	return req, nil
}

func decodeDonationsRequest(r *http.Request) (dto.ListDonationsRequest, error) {
	value := r.URL.Query()

	offset, err := strconv.Atoi(value.Get("offset"))
	if err != nil || offset < 0 {
		return dto.ListDonationsRequest{}, internal_errors.BadRequest{Message: "Invalid offset value"}
	}

	limit, err := strconv.Atoi(value.Get("limit"))
	if err != nil || limit < 0 {
		return dto.ListDonationsRequest{}, internal_errors.BadRequest{Message: "Invalid limit value"}
	}

	// Keeping default as ascending order
	isAscending, err := strconv.ParseBool(value.Get("is_ascending"))
	if err != nil {
		isAscending = true
	}

	req := dto.ListDonationsRequest{
		Search:             value.Get("search"),
		IsAnonymous:        value.Get("is_anonymous"),
		Offset:             uint(offset),
		Limit:              uint(limit),
		OrderByKey:         value.Get("order_by"),
		OrderByIsAscending: isAscending,
	}
	return req, nil
}
