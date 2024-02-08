package api

import (
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/fundraiser"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
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

		err = fundSvc.CreateFundraiser(req)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusCreated, dto.MessageResponse{
			Message: "Fundraiser created successfully",
		})
	}
}
