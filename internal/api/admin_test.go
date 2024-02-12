package api

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/admin/mocks"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/stretchr/testify/mock"
)

func TestLoginAdminHandler(t *testing.T) {
	adminSvc := mocks.NewService(t)
	userLoginHandler := LoginAdminHandler(adminSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for user login",
			input: `{
						"username": "admin",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginAdmin", mock.Anything).Return("token", nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail Invalid json",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail username not found",
			input: `{
						"username": "admin"
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail password not found",
			input: `{
						"password": "1"
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail username and password invalid",
			input: `{
						"username": "",
						"password": ""
					}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail login admin failed",
			input: `{
						"username": "admin",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginAdmin", mock.Anything).Return("", internal_errors.NotFoundError{Message: "Not found"}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "Fail login admin failed",
			input: `{
						"username": "admin",
						"password": "123"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginAdmin", mock.Anything).Return("", errors.New("Error in login")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(adminSvc)

			req, err := http.NewRequest("POST", "/admin/login", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(userLoginHandler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
