package organizer

import (
	"errors"
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestRegisterOrganizer(t *testing.T) {
	orgRepo := mocks.NewOrganizerStorer(t)
	service := NewService(orgRepo)

	tests := []struct {
		name            string
		input           dto.RegisterOrganizerRequest
		setup           func(orgMock *mocks.OrganizerStorer)
		isErrorExpected bool
	}{
		{
			name: "Success register organizer",
			input: dto.RegisterOrganizerRequest{
				Name:     "org",
				Detail:   "detail",
				Email:    "detail@gmail.com",
				Password: "123",
				Mobile:   "9876543210",
			},
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("RegisterOrganizer", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail register organizer",
			input: dto.RegisterOrganizerRequest{
				Name:     "org",
				Detail:   "detail",
				Email:    "detail@gmail.com",
				Password: "123",
				Mobile:   "9876543210",
			},
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("RegisterOrganizer", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgRepo)

			// test service
			err := service.RegisterOrganizer(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestDeleteOrganizer(t *testing.T) {
	orgRepo := mocks.NewOrganizerStorer(t)
	service := NewService(orgRepo)

	tests := []struct {
		name            string
		input           uint
		setup           func(orgMock *mocks.OrganizerStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success delete organizer",
			input: 1,
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("DeleteOrganizer", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Fail delete organizer",
			input: 1,
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("DeleteOrganizer", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgRepo)

			// test service
			err := service.DeleteOrganizer(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestLoginOrganizer(t *testing.T) {
	orgRepo := mocks.NewOrganizerStorer(t)
	service := NewService(orgRepo)

	tests := []struct {
		name            string
		input           dto.LoginOrganizerRequest
		setup           func(orgMock *mocks.OrganizerStorer, password string)
		isErrorExpected bool
	}{
		{
			name: "Success for Organizer Login",
			input: dto.LoginOrganizerRequest{
				Email:    "org@gmail.com",
				Password: "123",
			},
			setup: func(orgMock *mocks.OrganizerStorer, password string) {
				hashedPassword, _ := helpers.HashPassword(password)
				orgMock.On("GetOrganizerIDPasswordAndVerifyStatus", mock.Anything).Return(uint(1), hashedPassword, true, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail for not match",
			input: dto.LoginOrganizerRequest{
				Email:    "org@gmail.com",
				Password: "123",
			},
			setup: func(orgMock *mocks.OrganizerStorer, _ string) {
				hashedPassword, _ := helpers.HashPassword("other")
				orgMock.On("GetOrganizerIDPasswordAndVerifyStatus", mock.Anything).Return(uint(1), hashedPassword, true, nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail for GetOrganizerIDPasswordAndVerifyStatus",
			input: dto.LoginOrganizerRequest{
				Email:    "org@gmail.com",
				Password: "123",
			},
			setup: func(orgMock *mocks.OrganizerStorer, _ string) {
				orgMock.On("GetOrganizerIDPasswordAndVerifyStatus", mock.Anything).Return(uint(0), "", false, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail isVerified",
			input: dto.LoginOrganizerRequest{
				Email:    "org@gmail.com",
				Password: "123",
			},
			setup: func(orgMock *mocks.OrganizerStorer, password string) {
				hashedPassword, _ := helpers.HashPassword(password)
				orgMock.On("GetOrganizerIDPasswordAndVerifyStatus", mock.Anything).Return(uint(1), hashedPassword, false, nil).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgRepo, test.input.Password)

			// test service
			_, err := service.LoginOrganizer(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestVerifyOrganizer(t *testing.T) {
	orgRepo := mocks.NewOrganizerStorer(t)
	service := NewService(orgRepo)

	tests := []struct {
		name            string
		input           uint
		setup           func(orgMock *mocks.OrganizerStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success verify organizer",
			input: 1,
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("VerifyOrganizer", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Fail verify organizer",
			input: 1,
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("VerifyOrganizer", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgRepo)

			// test service
			err := service.VerifyOrganizer(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestGetOrganizerList(t *testing.T) {
	orgRepo := mocks.NewOrganizerStorer(t)
	service := NewService(orgRepo)

	tests := []struct {
		name            string
		input           dto.ListOrganizersRequest
		setup           func(orgMock *mocks.OrganizerStorer)
		isErrorExpected bool
	}{
		{
			name: "Success all organizers",
			input: dto.ListOrganizersRequest{
				Search:             "",
				Verified:           "true",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "fundraiser_id",
				OrderByIsAscending: true,
			},
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("GetOrganizerListCount", mock.Anything).Return(uint(10), nil).Once()
				orgMock.On("GetOrganizerList", mock.Anything).Return([]dto.OrganizerView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail GetOrganizerListCount",
			input: dto.ListOrganizersRequest{
				Search:             "",
				Verified:           "true",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "fundraiser_id",
				OrderByIsAscending: true,
			},
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("GetOrganizerListCount", mock.Anything).Return(uint(0), errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail GetOrganizerList",
			input: dto.ListOrganizersRequest{
				Search:             "",
				Verified:           "true",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "fundraiser_id",
				OrderByIsAscending: true,
			},
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("GetOrganizerListCount", mock.Anything).Return(uint(0), nil).Once()
				orgMock.On("GetOrganizerList", mock.Anything).Return([]dto.OrganizerView{}, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgRepo)

			// test service
			_, _, err := service.GetOrganizerList(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestGetOrganizer(t *testing.T) {
	orgRepo := mocks.NewOrganizerStorer(t)
	service := NewService(orgRepo)

	tests := []struct {
		name            string
		input           uint
		setup           func(orgMock *mocks.OrganizerStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success get organizer",
			input: 1,
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("GetOrganizer", mock.Anything).Return(dto.OrganizerView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Fail get organizer",
			input: 1,
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("GetOrganizer", mock.Anything).Return(dto.OrganizerView{}, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgRepo)

			// test service
			_, err := service.GetOrganizer(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestUpdateOrganizer(t *testing.T) {
	orgRepo := mocks.NewOrganizerStorer(t)
	service := NewService(orgRepo)

	tests := []struct {
		name            string
		input           dto.UpdateOrganizerRequest
		setup           func(orgMock *mocks.OrganizerStorer)
		isErrorExpected bool
	}{
		{
			name: "Success update organizer",
			input: dto.UpdateOrganizerRequest{
				OrganizerId: 1,
				Detail:      "detail",
				Email:       "detail@gmail.com",
				Mobile:      "9876543210",
			},
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("UpdateOrganizer", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail update organizer",
			input: dto.UpdateOrganizerRequest{
				OrganizerId: 1,
				Detail:      "detail",
				Email:       "detail@gmail.com",
				Mobile:      "9876543210",
			},
			setup: func(orgMock *mocks.OrganizerStorer) {
				orgMock.On("UpdateOrganizer", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orgRepo)

			// test service
			err := service.UpdateOrganizer(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
