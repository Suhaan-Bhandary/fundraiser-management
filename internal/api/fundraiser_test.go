package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	donationMocks "github.com/Suhaan-Bhandary/fundraiser-management/internal/app/donation/mocks"
	fundMocks "github.com/Suhaan-Bhandary/fundraiser-management/internal/app/fundraiser/mocks"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func TestCreateFundraiserHandler(t *testing.T) {
	fundSvc := fundMocks.NewService(t)
	userLoginHandler := CreateFundraiserHandler(fundSvc)

	tests := []struct {
		name               string
		input              string
		token              dto.Token
		isTokenPresent     bool
		setup              func(mock *fundMocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for Create fundraiser",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("CreateFundraiser", mock.Anything).Return(uint(1), nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:  "Fail for invalid json",
			input: "",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for Create fundraiser, title not present",
			input: `
			{
				"description": "Example",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for Create fundraiser, description not present",
			input: `
			{
				"title": "Fundraiser",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for Create fundraiser, image url not present",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for Create fundraiser, video url not present",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"organizer_id": 1,
				"image_url": "image",
				"target_amount": 2000
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for Create fundraiser, amount not present",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"organizer_id": 1,
				"image_url": "image",
				"video_url": "image"
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for Create fundraiser, token invalid",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			token:              dto.Token{},
			isTokenPresent:     false,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name: "Fail when Create fundraiser is called",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("CreateFundraiser", mock.Anything).Return(uint(0), errors.New("Error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundSvc)

			req, err := http.NewRequest("POST", "/fundraiser", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			if test.isTokenPresent {
				req = req.WithContext(context.WithValue(req.Context(), constants.TokenKey, test.token))
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(userLoginHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestDeleteFundraiserHandler(t *testing.T) {
	userSvc := fundMocks.NewService(t)
	deleteFundraiserHandler := DeleteFundraiserHandler(userSvc)

	tests := []struct {
		name               string
		input              string
		token              dto.Token
		isTokenPresent     bool
		setup              func(mock *fundMocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Success for delete user",
			input: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ADMIN,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("DeleteFundraiser", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:  "Fail: Invalid id",
			input: "",
			token: dto.Token{
				ID:   1,
				Role: constants.ADMIN,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Fail: Invalid token data",
			input:              "1",
			token:              dto.Token{},
			isTokenPresent:     false,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:  "Fail: Delete Fundraiser fail",
			input: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ADMIN,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("DeleteFundraiser", mock.Anything).Return(errors.New("Errors")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:  "Fail: Delete Fundraiser fail, for invalid credential",
			input: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("DeleteFundraiser", mock.Anything).Return(
					internal_errors.InvalidCredentialError{Message: "Invalid credentials"},
				).Once()
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userSvc)

			req, err := http.NewRequest(
				"DELETE",
				fmt.Sprintf("/user/%s", test.input), bytes.NewBuffer([]byte("")),
			)
			if err != nil {
				t.Fatal(err)
				return
			}

			// Adding token data
			if test.isTokenPresent {
				req = req.WithContext(context.WithValue(req.Context(), constants.TokenKey, test.token))
			}

			// Adding url params
			req = mux.SetURLVars(req, map[string]string{
				"id": test.input,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(deleteFundraiserHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestCreateDonationHandler(t *testing.T) {
	donationSvc := donationMocks.NewService(t)
	userLoginHandler := CreateDonationHandler(donationSvc)

	tests := []struct {
		name               string
		input              string
		fundraiserId       string
		token              dto.Token
		isTokenPresent     bool
		setup              func(mock *donationMocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for Create donation",
			input: `
			{
				"amount": 1000,
				"is_anonymous": true
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.USER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *donationMocks.Service) {
				mockSvc.On("CreateDonation", mock.Anything).Return(uint(1), nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Success for Create donation without is_anonymous",
			input: `
			{
				"amount": 1000
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.USER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *donationMocks.Service) {
				mockSvc.On("CreateDonation", mock.Anything).Return(uint(1), nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:         "Fail invalid json",
			input:        "",
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.USER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *donationMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail fundraiser id not found",
			input: `
			{
				"amount": 1000,
				"is_anonymous": true
			}
			`,
			fundraiserId: "",
			token: dto.Token{
				ID:   1,
				Role: constants.USER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *donationMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail no token",
			input: `
			{
				"amount": 100
			}
			`,
			fundraiserId:       "1",
			token:              dto.Token{},
			isTokenPresent:     false,
			setup:              func(mockSvc *donationMocks.Service) {},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name: "Fail invalid amount",
			input: `
				{
					"amount": -100
				}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.USER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *donationMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail create donation",
			input: `
			{
				"amount": 1000,
				"is_anonymous": true
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.USER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *donationMocks.Service) {
				mockSvc.On("CreateDonation", mock.Anything).Return(uint(0), errors.New("Fail")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(donationSvc)

			req, err := http.NewRequest(
				"POST",
				fmt.Sprintf("/fundraiser/%s", test.fundraiserId), bytes.NewBuffer([]byte(test.input)),
			)
			if err != nil {
				t.Fatal(err)
				return
			}

			if test.isTokenPresent {
				req = req.WithContext(context.WithValue(req.Context(), constants.TokenKey, test.token))
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.fundraiserId,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(userLoginHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetFundraiserHandler(t *testing.T) {
	fundSvc := fundMocks.NewService(t)
	getFundraiserHandler := GetFundraiserHandler(fundSvc)

	tests := []struct {
		name               string
		fundraiserId       string
		setup              func(mock *fundMocks.Service)
		expectedStatusCode int
	}{
		{
			name:         "Success for Get fundraiser",
			fundraiserId: "1",
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("GetFundraiserDetail", mock.Anything).Return(dto.FundraiserView{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail invalid id",
			fundraiserId:       "name",
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:         "Fail as GetFundraiserDetail fails",
			fundraiserId: "1",
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("GetFundraiserDetail", mock.Anything).Return(dto.FundraiserView{}, errors.New("Error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:         "Fail as GetFundraiserDetail fails",
			fundraiserId: "100",
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("GetFundraiserDetail", mock.Anything).
					Return(dto.FundraiserView{}, internal_errors.NotFoundError{Message: "fundraiser Not found"}).
					Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundSvc)

			req, err := http.NewRequest(
				"GET",
				fmt.Sprintf("/fundraiser/%s", test.fundraiserId), bytes.NewBuffer([]byte("")),
			)
			if err != nil {
				t.Fatal(err)
				return
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.fundraiserId,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getFundraiserHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestCloseFundraiserHandler(t *testing.T) {
	fundSvc := fundMocks.NewService(t)
	closeFundraiserHandler := CloseFundraiserHandler(fundSvc)

	tests := []struct {
		name               string
		fundraiserId       string
		token              dto.Token
		isTokenPresent     bool
		setup              func(mock *fundMocks.Service)
		expectedStatusCode int
	}{
		{
			name:         "Success for close fundraiser",
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("CloseFundraiser", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:         "Fail incorrect fundraiser",
			fundraiserId: "name",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Fail token not present",
			fundraiserId:       "1",
			token:              dto.Token{},
			isTokenPresent:     false,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:         "Fail CloseFundraiser",
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("CloseFundraiser", mock.Anything, mock.Anything).Return(errors.New("Error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:         "Fail fundraiser not found",
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("CloseFundraiser", mock.Anything, mock.Anything).Return(internal_errors.NotFoundError{Message: "Fundraiser not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundSvc)

			req, err := http.NewRequest(
				"PATCH",
				fmt.Sprintf("/fundraiser/%s/close", test.fundraiserId), bytes.NewBuffer([]byte("")),
			)
			if err != nil {
				t.Fatal(err)
				return
			}

			if test.isTokenPresent {
				req = req.WithContext(context.WithValue(req.Context(), constants.TokenKey, test.token))
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.fundraiserId,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(closeFundraiserHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestBanFundraiserHandler(t *testing.T) {
	fundSvc := fundMocks.NewService(t)
	banFundraiserHandler := BanFundraiserHandler(fundSvc)

	tests := []struct {
		name               string
		fundraiserId       string
		setup              func(mock *fundMocks.Service)
		expectedStatusCode int
	}{
		{
			name:         "Success for ban fundraiser",
			fundraiserId: "1",
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("BanFundraiser", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail incorrect fundraiser",
			fundraiserId:       "name",
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:         "Fail ban fundraiser",
			fundraiserId: "1",
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("BanFundraiser", mock.Anything, mock.Anything).Return(errors.New("Error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:         "Fail fundraiser not found",
			fundraiserId: "1",
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("BanFundraiser", mock.Anything, mock.Anything).Return(internal_errors.NotFoundError{Message: "Fundraiser not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundSvc)

			req, err := http.NewRequest(
				"PATCH",
				fmt.Sprintf("/fundraiser/%s/ban", test.fundraiserId), bytes.NewBuffer([]byte("")),
			)
			if err != nil {
				t.Fatal(err)
				return
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.fundraiserId,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(banFundraiserHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestUnBanFundraiserHandler(t *testing.T) {
	fundSvc := fundMocks.NewService(t)
	unBanFundraiserHandler := UnBanFundraiserHandler(fundSvc)

	tests := []struct {
		name               string
		fundraiserId       string
		setup              func(mock *fundMocks.Service)
		expectedStatusCode int
	}{
		{
			name:         "Success for unban fundraiser",
			fundraiserId: "1",
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("UnBanFundraiser", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail incorrect fundraiser",
			fundraiserId:       "name",
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:         "Fail unban fundraiser",
			fundraiserId: "1",
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("UnBanFundraiser", mock.Anything, mock.Anything).Return(errors.New("Error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:         "Fail fundraiser not found",
			fundraiserId: "1",
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("UnBanFundraiser", mock.Anything, mock.Anything).Return(internal_errors.NotFoundError{Message: "Fundraiser not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundSvc)

			req, err := http.NewRequest(
				"PATCH",
				fmt.Sprintf("/fundraiser/%s/unban", test.fundraiserId), bytes.NewBuffer([]byte("")),
			)
			if err != nil {
				t.Fatal(err)
				return
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.fundraiserId,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(unBanFundraiserHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestListFundraisersHandler(t *testing.T) {}

func TestListFundraiserDonationsHandler(t *testing.T) {}

func TestListDonationsHandler(t *testing.T) {}

func TestUpdateFundraiserHandler(t *testing.T) {
	fundSvc := fundMocks.NewService(t)
	updateFundraiserHandler := UpdateFundraiserHandler(fundSvc)

	tests := []struct {
		name               string
		input              string
		fundraiserId       string
		token              dto.Token
		isTokenPresent     bool
		setup              func(mock *fundMocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for Update fundraiser",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("UpdateFundraiser", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:         "Fail invalid json",
			input:        "",
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail no token passed",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			fundraiserId:       "1",
			token:              dto.Token{},
			isTokenPresent:     false,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name: "Fail title not found",
			input: `
			{
				"description": "Example",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail description not found",
			input: `
			{
				"title": "Fundraiser",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail image url not found",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail video url not found",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail target amount not found",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"image_url": "image",
				"video_url": "image"
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *fundMocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail update fundraiser",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("UpdateFundraiser", mock.Anything).Return(errors.New("Error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Fail fundraiser not found",
			input: `
			{
				"title": "Fundraiser",
				"description": "Example",
				"image_url": "image",
				"video_url": "image",
				"target_amount": 2000
			}
			`,
			fundraiserId: "1",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *fundMocks.Service) {
				mockSvc.On("UpdateFundraiser", mock.Anything).Return(internal_errors.NotFoundError{Message: "Not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundSvc)

			req, err := http.NewRequest("PUT", "/fundraiser/id", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			if test.isTokenPresent {
				req = req.WithContext(context.WithValue(req.Context(), constants.TokenKey, test.token))
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.fundraiserId,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(updateFundraiserHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
