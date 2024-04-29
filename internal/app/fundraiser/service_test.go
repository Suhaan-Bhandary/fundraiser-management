package fundraiser

import (
	"errors"
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCreateFundraiser(t *testing.T) {
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(fundRepo)

	tests := []struct {
		name            string
		input           dto.CreateFundraiserRequest
		setup           func(fundMock *mocks.FundraiserStorer)
		isErrorExpected bool
	}{
		{
			name: "Success for create fundraiser",
			input: dto.CreateFundraiserRequest{
				Title:        "Fundraiser",
				Description:  "Description",
				OrganizerId:  1,
				ImageUrl:     "image",
				VideoUrl:     "video",
				TargetAmount: 1000,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("CreateFundraiser", mock.Anything).Return(uint(1), nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail create fundraiser",
			input: dto.CreateFundraiserRequest{
				Title:        "Fundraiser",
				Description:  "Description",
				OrganizerId:  1,
				ImageUrl:     "image",
				VideoUrl:     "video",
				TargetAmount: 1000,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("CreateFundraiser", mock.Anything).Return(uint(0), errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundRepo)

			// test service
			_, err := service.CreateFundraiser(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestDeleteFundraiser(t *testing.T) {
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(fundRepo)

	tests := []struct {
		name            string
		input           dto.DeleteFundraiserRequest
		setup           func(fundMock *mocks.FundraiserStorer)
		isErrorExpected bool
	}{
		{
			name: "Success for organizer delete fundraiser",
			input: dto.DeleteFundraiserRequest{
				Token: dto.Token{
					ID:   1,
					Role: constants.ORGANIZER,
				},
				FundraiserId: 1,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerId", mock.Anything).Return(uint(1), nil).Once()
				fundMock.On("DeleteFundraiser", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Success for admin delete fundraiser",
			input: dto.DeleteFundraiserRequest{
				Token: dto.Token{
					ID:   1,
					Role: constants.ADMIN,
				},
				FundraiserId: 1,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("DeleteFundraiser", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Error for get organizer id",
			input: dto.DeleteFundraiserRequest{
				Token: dto.Token{
					ID:   1,
					Role: constants.ORGANIZER,
				},
				FundraiserId: 1,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerId", mock.Anything).Return(uint(0), errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Error for get creator and organizer id mismatch",
			input: dto.DeleteFundraiserRequest{
				Token: dto.Token{
					ID:   1,
					Role: constants.ORGANIZER,
				},
				FundraiserId: 1,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerId", mock.Anything).Return(uint(10), nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Error for delete fundraiser",
			input: dto.DeleteFundraiserRequest{
				Token: dto.Token{
					ID:   1,
					Role: constants.ADMIN,
				},
				FundraiserId: 1,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("DeleteFundraiser", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundRepo)

			// test service
			err := service.DeleteFundraiser(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestGetFundraiserDetail(t *testing.T) {
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(fundRepo)

	tests := []struct {
		name            string
		input           uint
		setup           func(fundMock *mocks.FundraiserStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success for organizer delete fundraiser",
			input: 1,
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiser", mock.Anything).Return(dto.FundraiserView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Error in GetFundraiser",
			input: 1,
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiser", mock.Anything).Return(dto.FundraiserView{}, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundRepo)

			// test service
			_, err := service.GetFundraiserDetail(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestCloseFundraiser(t *testing.T) {
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(fundRepo)

	tests := []struct {
		name            string
		fundraiserId    uint
		token           dto.Token
		setup           func(fundMock *mocks.FundraiserStorer)
		isErrorExpected bool
	}{
		{
			name:         "Success for close fundraiser",
			fundraiserId: 1,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerIdAndStatus", mock.Anything).Return(uint(1), constants.ACTIVE_STATUS, nil).Once()
				fundMock.On("CloseFundraiser", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name:         "Fail GetFundraiserOrganizerIdAndStatus",
			fundraiserId: 1,
			token: dto.Token{
				ID:   99,
				Role: constants.ORGANIZER,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerIdAndStatus", mock.Anything).Return(uint(0), "", errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name:         "Fail for different organizer",
			fundraiserId: 1,
			token: dto.Token{
				ID:   99,
				Role: constants.ORGANIZER,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerIdAndStatus", mock.Anything).Return(uint(1), constants.ACTIVE_STATUS, nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name:         "Fail for status inactive",
			fundraiserId: 1,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerIdAndStatus", mock.Anything).Return(uint(1), constants.INACTIVE_STATUS, nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name:         "Fail for status banned",
			fundraiserId: 1,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerIdAndStatus", mock.Anything).Return(uint(1), constants.BANNED_STATUS, nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name:         "Fail CloseFundraiser",
			fundraiserId: 1,
			token: dto.Token{
				ID:   1,
				Role: constants.ORGANIZER,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerIdAndStatus", mock.Anything).Return(uint(1), constants.ACTIVE_STATUS, nil).Once()
				fundMock.On("CloseFundraiser", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundRepo)

			// test service
			err := service.CloseFundraiser(test.fundraiserId, test.token)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestBanFundraiser(t *testing.T) {
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(fundRepo)

	tests := []struct {
		name            string
		input           uint
		setup           func(fundMock *mocks.FundraiserStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success for ban fundraiser",
			input: 1,
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("BanFundraiser", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Error in BanFundraiser",
			input: 1,
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("BanFundraiser", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundRepo)

			// test service
			err := service.BanFundraiser(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestUnBanFundraiser(t *testing.T) {
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(fundRepo)

	tests := []struct {
		name            string
		input           uint
		setup           func(fundMock *mocks.FundraiserStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success for unban fundraiser",
			input: 1,
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("UnBanFundraiser", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Error in UnBanFundraiser",
			input: 1,
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("UnBanFundraiser", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundRepo)

			// test service
			err := service.UnBanFundraiser(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestListFundraisers(t *testing.T) {
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(fundRepo)

	tests := []struct {
		name            string
		input           dto.ListFundraisersRequest
		setup           func(fundMock *mocks.FundraiserStorer)
		isErrorExpected bool
	}{
		{
			name: "Success all fundraisers",
			input: dto.ListFundraisersRequest{
				Search:             "donation",
				Status:             constants.ACTIVE_STATUS,
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetListFundraisersCount", mock.Anything).Return(uint(10), nil).Once()
				fundMock.On("ListFundraiser", mock.Anything).Return([]dto.FundraiserView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail GetListFundraisersCount",
			input: dto.ListFundraisersRequest{
				Search:             "donation",
				Status:             constants.ACTIVE_STATUS,
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetListFundraisersCount", mock.Anything).Return(uint(0), errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail ListFundraiser",
			input: dto.ListFundraisersRequest{
				Search:             "donation",
				Status:             constants.ACTIVE_STATUS,
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetListFundraisersCount", mock.Anything).Return(uint(10), nil).Once()
				fundMock.On("ListFundraiser", mock.Anything).Return([]dto.FundraiserView{}, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundRepo)

			// test service
			_, _, err := service.ListFundraisers(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestUpdateFundraiser(t *testing.T) {
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(fundRepo)

	tests := []struct {
		name            string
		input           dto.UpdateFundraiserRequest
		setup           func(fundMock *mocks.FundraiserStorer)
		isErrorExpected bool
	}{
		{
			name: "Success for update fundraiser",
			input: dto.UpdateFundraiserRequest{
				RequestOrganizerId: 1,
				FundraiserId:       1,
				Title:              "Fundraiser",
				Description:        "Description",
				ImageUrl:           "image",
				VideoUrl:           "video",
				TargetAmount:       1000,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerId", mock.Anything).Return(uint(1), nil).Once()
				fundMock.On("UpdateFundraiser", mock.Anything).Return(nil).Once()
				fundMock.On("UpdateFundraiserStatus", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail as organizer doesnot match",
			input: dto.UpdateFundraiserRequest{
				RequestOrganizerId: 1,
				FundraiserId:       1,
				Title:              "Fundraiser",
				Description:        "Description",
				ImageUrl:           "image",
				VideoUrl:           "video",
				TargetAmount:       1000,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerId", mock.Anything).Return(uint(10), nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail GetFundraiserOrganizerId",
			input: dto.UpdateFundraiserRequest{
				RequestOrganizerId: 1,
				FundraiserId:       1,
				Title:              "Fundraiser",
				Description:        "Description",
				ImageUrl:           "image",
				VideoUrl:           "video",
				TargetAmount:       1000,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerId", mock.Anything).Return(uint(0), errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail UpdateFundraiser",
			input: dto.UpdateFundraiserRequest{
				RequestOrganizerId: 1,
				FundraiserId:       1,
				Title:              "Fundraiser",
				Description:        "Description",
				ImageUrl:           "image",
				VideoUrl:           "video",
				TargetAmount:       1000,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerId", mock.Anything).Return(uint(1), nil).Once()
				fundMock.On("UpdateFundraiser", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail for update fundraiser status",
			input: dto.UpdateFundraiserRequest{
				RequestOrganizerId: 1,
				FundraiserId:       1,
				Title:              "Fundraiser",
				Description:        "Description",
				ImageUrl:           "image",
				VideoUrl:           "video",
				TargetAmount:       1000,
			},
			setup: func(fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserOrganizerId", mock.Anything).Return(uint(1), nil).Once()
				fundMock.On("UpdateFundraiser", mock.Anything).Return(nil).Once()
				fundMock.On("UpdateFundraiserStatus", mock.Anything).Return(errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(fundRepo)

			// test service
			err := service.UpdateFundraiser(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
