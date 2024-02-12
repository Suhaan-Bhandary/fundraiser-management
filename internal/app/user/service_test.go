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
				userMock.On("GetUserIDPassword", mock.Anything).Return(uint(1), hashedPassword, nil).Once()
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
				userMock.On("GetUserIDPassword", mock.Anything).Return(uint(0), "", errors.New("Error")).Once()
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
				userMock.On("GetUserIDPassword", mock.Anything).Return(uint(0), hashedPassword, nil).Once()
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

func TestDeleteUser(t *testing.T) {
	userRepo := mocks.NewUserStorer(t)
	service := NewService(userRepo)

	tests := []struct {
		name            string
		input           uint
		setup           func(userMock *mocks.UserStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success delete user",
			input: 1,
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("DeleteUser", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Fail delete user",
			input: 1,
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("DeleteUser", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userRepo)

			// test service
			err := service.DeleteUser(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestListUsers(t *testing.T) {
	userRepo := mocks.NewUserStorer(t)
	service := NewService(userRepo)

	tests := []struct {
		name            string
		input           dto.ListUserRequest
		setup           func(userMock *mocks.UserStorer)
		isErrorExpected bool
	}{
		{
			name: "Success all users",
			input: dto.ListUserRequest{
				Search:             "search",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("GetListUsersCount", mock.Anything).Return(uint(10), nil).Once()
				userMock.On("ListUsers", mock.Anything).Return([]dto.UserView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail GetListUsersCount",
			input: dto.ListUserRequest{
				Search:             "search",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("GetListUsersCount", mock.Anything).Return(uint(0), errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail ListUsers",
			input: dto.ListUserRequest{
				Search:             "search",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("GetListUsersCount", mock.Anything).Return(uint(10), nil).Once()
				userMock.On("ListUsers", mock.Anything).Return([]dto.UserView{}, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userRepo)

			// test service
			_, _, err := service.ListUsers(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestGetUserProfile(t *testing.T) {
	userRepo := mocks.NewUserStorer(t)
	service := NewService(userRepo)

	tests := []struct {
		name            string
		userId          uint
		setup           func(userMock *mocks.UserStorer)
		isErrorExpected bool
	}{
		{
			name:   "Success get user",
			userId: 1,
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("GetUserProfile", mock.Anything).Return(dto.UserView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name:   "Fail get user",
			userId: 1,
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("GetUserProfile", mock.Anything).Return(dto.UserView{}, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userRepo)

			// test service
			_, err := service.GetUserProfile(test.userId)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
