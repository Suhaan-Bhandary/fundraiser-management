package api

import (
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/donation"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/fundraiser"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/middleware"
)

func CreateFundraiserHandler(fundSvc fundraiser.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeCreateFundraiser(r)
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
		req, err := decodeDeleteFundraiser(r)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		err = fundSvc.DeleteFundraiser(req)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, dto.MessageResponse{
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
		fmt.Println(donationId, err)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusCreated, dto.CreateDonationResponse{
			DonationId: donationId,
		})
	}
}

func GetFundraiserHandler(fundSvc fundraiser.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fundraiserId, err := decodeId(r)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		fundraiserDetail, err := fundSvc.GetFundraiserDetail(fundraiserId)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, dto.GetFundraiserResponse{
			Fundraiser: fundraiserDetail,
		})
	}
}

func CloseFundraiserHandler(fundSvc fundraiser.Service) func(http.ResponseWriter, *http.Request) {
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

		err = fundSvc.CloseFundraiser(fundraiserId, tokenData)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, dto.MessageResponse{
			Message: "Fundraiser closed successfully",
		})
	}
}

func BanFundraiserHandler(fundSvc fundraiser.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fundraiserId, err := decodeId(r)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = fundSvc.BanFundraiser(fundraiserId)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, dto.MessageResponse{
			Message: "Fundraiser Banned successfully",
		})
	}
}

func UnBanFundraiserHandler(fundSvc fundraiser.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fundraiserId, err := decodeId(r)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = fundSvc.UnBanFundraiser(fundraiserId)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, dto.MessageResponse{
			Message: "Fundraiser UnBanned successfully",
		})
	}
}

func ListFundraisersHandler(fundSvc fundraiser.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeListFundraisersRequest(r)
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

		fundraisers, totalCount, err := fundSvc.ListFundraisers(req)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, dto.ListFundraisersResponse{
			Fundraisers: fundraisers,
			TotalCount:  totalCount,
		})
	}
}

func ListFundraiserDonationsHandler(donationSvc donation.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeFundraiserDonationsRequest(r)
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

		donations, totalCount, err := donationSvc.ListFundraiserDonations(req)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, dto.ListFundraiserDonationsResponse{
			Donations:  donations,
			TotalCount: totalCount,
		})
	}
}

func ListDonationsHandler(donationSvc donation.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeDonationsRequest(r)
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

		donations, totalCount, err := donationSvc.ListDonations(req)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, dto.ListDonationsResponse{
			Donations:  donations,
			TotalCount: totalCount,
		})
	}
}

func UpdateFundraiserHandler(fundraiserSvc fundraiser.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeUpdateFundraiser(r)
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

		err = fundraiserSvc.UpdateFundraiser(req)
		if err != nil {
			statusCode, errResponse := internal_errors.MatchError(err)
			middleware.ErrorResponse(w, statusCode, errResponse)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, dto.MessageResponse{
			Message: "Updated fundraiser successfully",
		})
	}
}
