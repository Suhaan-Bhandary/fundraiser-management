package admin

import (
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestLoginAdmin(t *testing.T) {
	adminRepo := mocks.NewAdminStorer(t)
	service := NewService(adminRepo)

	tests := []struct {
		name            string
		input           dto.LoginAdminRequest
		setup           func(userMock *mocks.AdminStorer, password string)
		isErrorExpected bool
	}{
		{
			name: "Success for admin login",
			input: dto.LoginAdminRequest{
				Username: "test",
				Password: "123",
			},
			setup: func(userMock *mocks.AdminStorer, password string) {
				hashedPassword, _ := helpers.HashPassword(password)
				userMock.On("GetAdminIDPassword", mock.Anything).Return(uint(1), hashedPassword, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Failed because admin not found",
			input: dto.LoginAdminRequest{
				Username: "test",
				Password: "123",
			},
			setup: func(userMock *mocks.AdminStorer, _ string) {
				userMock.On("GetAdminIDPassword", mock.Anything).Return(uint(0), "", internal_errors.NotFoundError{Message: "Admin not found"}).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Failed because invalid password",
			input: dto.LoginAdminRequest{
				Username: "test",
				Password: "123",
			},
			setup: func(userMock *mocks.AdminStorer, _ string) {
				hashedPassword, _ := helpers.HashPassword("other")
				userMock.On("GetAdminIDPassword", mock.Anything).Return(uint(0), hashedPassword, nil).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(adminRepo, test.input.Password)

			// test service
			_, err := service.LoginAdmin(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
