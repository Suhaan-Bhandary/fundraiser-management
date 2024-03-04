package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/user/mocks"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUserHandler(t *testing.T) {
	userSvc := mocks.NewService(t)
	userRegisterHandler := RegisterUserHandler(userSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for user Detail",
			input: `{
						"first_name": "suhaan",
						"last_name": "bhandary",
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterUser", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "Fail for incorrect json",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for missing first_name field",
			input: `{
						"last_name": "bhandary",
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for missing last_name field",
			input: `{
						"first_name": "suhaan",
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for missing email field",
			input: `{
						"first_name": "suhaan",
						"last_name": "bhandary",
						"password": "123"   
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for user password",
			input: `{
						"first_name": "suhaan",
						"last_name": "bhandary",
						"email": "suhaanbhandary1@gmail.com",
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Success for user Detail",
			input: `{
						"first_name": "suhaan",
						"last_name": "bhandary",
						"email": "suhaanbhandary1.com",
						"password": "123"   
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail from service",
			input: `{
						"first_name": "suhaan",
						"last_name": "bhandary",
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterUser", mock.Anything).Return(errors.New("Error from DB")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Fail from service because of duplicate email",
			input: `{
						"first_name": "suhaan",
						"last_name": "bhandary",
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterUser", mock.Anything).Return(internal_errors.DuplicateKeyError{Message: "duplicate email"}).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userSvc)

			req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(userRegisterHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestLoginUserHandler(t *testing.T) {
	userSvc := mocks.NewService(t)
	userLoginHandler := LoginUserHandler(userSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for user login",
			input: `{
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginUser", mock.Anything).Return("token", nil).Once()
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
			name: "Fail Error in LoginUser",
			input: `{
						"email": "suhaanbhandary1@gmail.com",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginUser", mock.Anything).Return("", errors.New("Error when login")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userSvc)

			req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
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

func TestDeleteUserHandler(t *testing.T) {
	userSvc := mocks.NewService(t)
	deleteUserHandler := DeleteUserHandler(userSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Success for delete user",
			input: "1",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteUser", mock.Anything).Return(nil).Once()
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
			name:  "Fail for delete user not found",
			input: "100",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteUser", mock.Anything).Return(internal_errors.NotFoundError{Message: "User not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:  "Fail for delete fail",
			input: "1",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteUser", mock.Anything).Return(
					errors.New("error"),
				).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
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

			req = mux.SetURLVars(req, map[string]string{
				"id": test.input,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(deleteUserHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestListUsersHandler(t *testing.T) {
	// TODO: Remaining
}

func TestGetUserProfileHandler(t *testing.T) {
	userSvc := mocks.NewService(t)
	getUserProfileHandler := GetUserProfileHandler(userSvc)

	tests := []struct {
		name               string
		token              dto.Token
		isTokenPresent     bool
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for user profile",
			token: dto.Token{
				ID:   1,
				Role: constants.USER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetUserProfile", mock.Anything).Return(dto.UserView{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Error for no token",
			token:              dto.Token{},
			isTokenPresent:     false,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name: "Fail as error in GetUserProfile",
			token: dto.Token{
				ID:   1,
				Role: constants.USER,
			},
			isTokenPresent: true,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetUserProfile", mock.Anything).Return(dto.UserView{}, errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userSvc)

			req, err := http.NewRequest("GET", "/user/profile", bytes.NewBuffer([]byte("")))
			if err != nil {
				t.Fatal(err)
				return
			}

			if test.isTokenPresent {
				req = req.WithContext(context.WithValue(req.Context(), constants.TokenKey, test.token))
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getUserProfileHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestListUserDonationsHandler(t *testing.T) {

}
