package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/user/mocks"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
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

			req, err := http.NewRequest("GET", "/user/register", bytes.NewBuffer([]byte(test.input)))
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
