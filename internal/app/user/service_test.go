package user

import (
	"errors"
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
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
				Password:  "hi",
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
				Password:  "hi",
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
