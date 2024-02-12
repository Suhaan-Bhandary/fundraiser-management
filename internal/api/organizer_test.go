package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/organizer/mocks"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func TestRegisterOrganizerHandler(t *testing.T) {
	orgSvc := mocks.NewService(t)
	organizerRegisterHandler := RegisterOrganizerHandler(orgSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for register organizer",
			input: `{
						"name": "test",
						"detail": "this is test detail",
						"email": "test@gmail.com",
						"password": "test",
						"mobile": "1234567890"
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterOrganizer", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "Fail invalid json",
			input:              ``,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail name not found",
			input: `{
						"detail": "this is test detail",
						"email": "test@gmail.com",
						"password": "test",
						"mobile": "1234567890"
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail detail not found",
			input: `{
						"name": "test",
						"email": "test@gmail.com",
						"password": "test",
						"mobile": "1234567890"
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail email not found",
			input: `{
						"name": "test",
						"detail": "this is test detail",
						"password": "test",
						"mobile": "1234567890"
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail invalid email",
			input: `{
						"name": "test",
						"email": "test.com",
						"detail": "this is test detail",
						"password": "test",
						"mobile": "1234567890"
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail password not found",
			input: `{
						"name": "test",
						"detail": "this is test detail",
						"email": "test@gmail.com",
						"mobile": "1234567890"
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail mobile not found",
			input: `{
						"name": "test",
						"detail": "this is test detail",
						"email": "test@gmail.com",
						"password": "test"
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Organizer register error",
			input: `{
						"name": "test",
						"detail": "this is test detail",
						"email": "test@gmail.com",
						"password": "test",
						"mobile": "1234567890"
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterOrganizer", mock.Anything).Return(errors.New("Error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Organizer email already present",
			input: `{
						"name": "test",
						"detail": "this is test detail",
						"email": "test@gmail.com",
						"password": "test",
						"mobile": "1234567890"
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterOrganizer", mock.Anything).Return(internal_errors.BadRequest{Message: "Email already present"}).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgSvc)

			req, err := http.NewRequest("POST", "/organizer/register", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(organizerRegisterHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestLoginOrganizerHandler(t *testing.T) {
	orgSvc := mocks.NewService(t)
	orgLoginHandler := LoginOrganizerHandler(orgSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for org login",
			input: `{
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginOrganizer", mock.Anything).Return("token", nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail for incorrect json",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for missing email field",
			input: `{
						"password": "123"   
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for incorrect email",
			input: `{
						"email": "suhaanbhandary1.com",
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for missing password field",
			input: `{
						"email": "suhaanbhandary1@gmail.com",
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail LoginOrganizer",
			input: `{
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginOrganizer", mock.Anything).Return("", errors.New("Error when login")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Fail email not found",
			input: `{
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginOrganizer", mock.Anything).Return("", internal_errors.NotFoundError{Message: "Not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgSvc)

			req, err := http.NewRequest("POST", "/organizer/login", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(orgLoginHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestDeleteOrganizerHandler(t *testing.T) {
	orgSvc := mocks.NewService(t)
	deleteOrgHandler := DeleteOrganizerHandler(orgSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Success for delete organizer",
			input: "1",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteOrganizer", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail for no id in url",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Fail for incorrect id in url",
			input:              "hi",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "Fail for delete organizer not found",
			input: "100",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteOrganizer", mock.Anything).Return(internal_errors.NotFoundError{Message: "User not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:  "Fail for DeleteOrganizer",
			input: "1",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteOrganizer", mock.Anything).Return(
					errors.New("error"),
				).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgSvc)

			req, err := http.NewRequest(
				"DELETE",
				fmt.Sprintf("/organizer/%s", test.input), bytes.NewBuffer([]byte("")),
			)
			if err != nil {
				t.Fatal(err)
				return
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.input,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(deleteOrgHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestVerifyOrganizerHandler(t *testing.T) {
	orgSvc := mocks.NewService(t)
	verifyOrganizerHandler := VerifyOrganizerHandler(orgSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Success for verify organizer",
			input: "1",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("VerifyOrganizer", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail for no id in url",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Fail for incorrect id in url",
			input:              "hi",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "Fail for organizer not found",
			input: "100",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("VerifyOrganizer", mock.Anything).Return(internal_errors.NotFoundError{Message: "organizer not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:  "Fail for VerifyOrganizer",
			input: "1",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("VerifyOrganizer", mock.Anything).Return(
					errors.New("error"),
				).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgSvc)

			req, err := http.NewRequest(
				"patch",
				fmt.Sprintf("/organizer/%s/verify", test.input), bytes.NewBuffer([]byte("")),
			)
			if err != nil {
				t.Fatal(err)
				return
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.input,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(verifyOrganizerHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestListOrganizersHandler(t *testing.T) {}

func TestGetOrganizerHandler(t *testing.T) {
	orgSvc := mocks.NewService(t)
	getOrganizerHandler := GetOrganizerHandler(orgSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Success for get organizer",
			input: "1",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetOrganizer", mock.Anything).Return(dto.OrganizerView{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail invalid id",
			input:              "name",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Fail no id param",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "Fail as error in GetOrganizer",
			input: "1",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetOrganizer", mock.Anything).Return(dto.OrganizerView{}, errors.New("Error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:  "Fail as organizer not found",
			input: "100",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetOrganizer", mock.Anything).Return(dto.OrganizerView{}, internal_errors.NotFoundError{Message: "Not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgSvc)

			req, err := http.NewRequest(
				"GET",
				fmt.Sprintf("/organizer/%s/verify", test.input), bytes.NewBuffer([]byte("")),
			)
			if err != nil {
				t.Fatal(err)
				return
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.input,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getOrganizerHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestUpdateOrganizerHandler(t *testing.T) {
	orgSvc := mocks.NewService(t)
	updateOrganizerHandler := UpdateOrganizerHandler(orgSvc)

	tests := []struct {
		name               string
		input              string
		token              dto.Token
		isTokenPresent     bool
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for Update organizer",
			input: `
			{
				"detail": "detail",
				"email": "test@gmail.com",
				"mobile": "1234567890"
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("UpdateOrganizer", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:  "Fail invalid json",
			input: "",
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail no token passed",
			input: `
			{
				"detail": "detail",
				"email": "test@gmail.com",
				"mobile": "1234567890"
			}
			`,
			token:              dto.Token{},
			isTokenPresent:     false,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name: "Fail detail not found",
			input: `
			{
				"email": "test@gmail.com",
				"mobile": "1234567890"
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail email not found",
			input: `
			{
				"detail": "detail",
				"mobile": "1234567890"
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid email not found",
			input: `
			{
				"email": "test.com",
				"detail": "detail",
				"mobile": "1234567890"
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail mobile not found",
			input: `
			{
				"email": "test.com",
				"detail": "detail"
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent:     true,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail update organizer",
			input: `
			{
				"detail": "test",
				"email": "test@gmail.com",
				"mobile": "1234567890"
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("UpdateOrganizer", mock.Anything).Return(errors.New("Error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Fail organizer not found",
			input: `
			{
				"detail": "test",
				"email": "test@gmail.com",
				"mobile": "1234567890"
			}
			`,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("UpdateOrganizer", mock.Anything).Return(internal_errors.NotFoundError{Message: "Not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgSvc)

			req, err := http.NewRequest("PUT", "/organizer", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			if test.isTokenPresent {
				req = req.WithContext(context.WithValue(req.Context(), constants.TokenKey, test.token))
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(updateOrganizerHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
