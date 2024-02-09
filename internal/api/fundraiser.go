package api

import (
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/donation"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/fundraiser"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/middleware"
)

func CreateFundraiserHandler(fundSvc fundraiser.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenData, err := decodeTokenFromContext(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		req, err := decodeCreateFundraiser(r)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		// Assigning the organizer id before validating it
		req.OrganizerId = uint(tokenData.ID)
		err = req.Validate()
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		fundraiserId, err := fundSvc.CreateFundraiser(req)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusCreated, dto.CreateFundraiserResponse{
			FundraiserId: fundraiserId,
		})
	}
}

func DeleteFundraiserHandler(fundSvc fundraiser.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fundraiserId, err := decodeId(r)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		tokenData, err := decodeTokenFromContext(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		err = fundSvc.DeleteFundraiser(fundraiserId, tokenData)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusCreated, dto.MessageResponse{
			Message: "Fundraiser deleted successfully",
		})
	}
}

func CreateDonationHandler(donationSvc donation.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeCreateDonation(r)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		err = req.Validate()
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		donationId, err := donationSvc.CreateDonation(req)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusCreated, dto.CreateDonationResponse{
			DonationId: donationId,
		})
	}
}
