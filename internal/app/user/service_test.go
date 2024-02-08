package user

import (
	"errors"
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	userRepo := mocks.NewUserStorer(t)
	service := NewService(userRepo)

	tests := []struct {
		name            string
		input           dto.RegisterUserRequest
		setup           func(userMock *mocks.UserStorer)
		isErrorExpected bool
	}{
		{
			name: "Success for user Detail",
			input: dto.RegisterUserRequest{
				FirstName: "suhaan",
				LastName:  "bhandary",
				Email:     "suhaanbhandary1@gmail.com",
				Password:  "123",
			},
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("RegisterUser", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Failed because RegisterUser",
			input: dto.RegisterUserRequest{
				FirstName: "suhaan",
				LastName:  "bhandary",
				Email:     "suhaanbhandary1@gmail.com",
				Password:  "123",
			},
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("RegisterUser", mock.Anything).Return(errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userRepo)

			// test service
			err := service.RegisterUser(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	userRepo := mocks.NewUserStorer(t)
	service := NewService(userRepo)

	tests := []struct {
		name            string
		input           dto.LoginUserRequest
		setup           func(userMock *mocks.UserStorer, password string)
		isErrorExpected bool
	}{
		{
			name: "Success for User Login",
			input: dto.LoginUserRequest{
				Email:    "suhaanbhandary1@gmail.com",
				Password: "123",
			},
			setup: func(userMock *mocks.UserStorer, password string) {
				hashedPassword, _ := helpers.HashPassword(password)
				userMock.On("GetUserIDPassword", mock.Anything).Return(0, hashedPassword, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Failed because invalid email",
			input: dto.LoginUserRequest{
				Email:    "abc@gmail.com",
				Password: "123",
			},
			setup: func(userMock *mocks.UserStorer, _ string) {
				userMock.On("GetUserIDPassword", mock.Anything).Return(-1, "", errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Failed because invalid password",
			input: dto.LoginUserRequest{
				Email:    "abc@gmail.com",
				Password: "123",
			},
			setup: func(userMock *mocks.UserStorer, _ string) {
				hashedPassword, _ := helpers.HashPassword("other")
				userMock.On("GetUserIDPassword", mock.Anything).Return(0, hashedPassword, nil).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userRepo, test.input.Password)

			// test service
			_, err := service.LoginUser(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
