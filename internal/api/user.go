package api

import (
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/user"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/middleware"
)

func RegisterUserHandler(userSvc user.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeRegisterUserRequest(r)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = req.Validate()
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = userSvc.RegisterUser(req)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusCreated, dto.MessageResponse{
			Message: "User registered successfully",
		})
	}
}
