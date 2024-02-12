package donation

import (
	"errors"
	"testing"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCreateDonation(t *testing.T) {
	donationRepo := mocks.NewDonationStorer(t)
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(donationRepo, fundRepo)

	tests := []struct {
		name            string
		input           dto.CreateDonationRequest
		setup           func(donationMock *mocks.DonationStorer, fundMock *mocks.FundraiserStorer)
		isErrorExpected bool
	}{
		{
			name: "Success for create donation",
			input: dto.CreateDonationRequest{
				UserId:       1,
				FundraiserId: 1,
				Amount:       10,
				IsAnonymous:  true,
			},
			setup: func(donationMock *mocks.DonationStorer, fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserStatus", mock.Anything).Return(constants.ACTIVE_STATUS, nil).Once()
				donationMock.On("CreateDonation", mock.Anything).Return(uint(1), nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail inactive fundraiser",
			input: dto.CreateDonationRequest{
				UserId:       1,
				FundraiserId: 1,
				Amount:       10,
				IsAnonymous:  true,
			},
			setup: func(_ *mocks.DonationStorer, fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserStatus", mock.Anything).Return(constants.INACTIVE_STATUS, nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail banned fundraiser",
			input: dto.CreateDonationRequest{
				UserId:       1,
				FundraiserId: 1,
				Amount:       10,
				IsAnonymous:  true,
			},
			setup: func(_ *mocks.DonationStorer, fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserStatus", mock.Anything).Return(constants.BANNED_STATUS, nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail fundraiser status",
			input: dto.CreateDonationRequest{
				UserId:       1,
				FundraiserId: 1,
				Amount:       10,
				IsAnonymous:  true,
			},
			setup: func(_ *mocks.DonationStorer, fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserStatus", mock.Anything).Return("", errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail error in create donation",
			input: dto.CreateDonationRequest{
				UserId:       1,
				FundraiserId: 1,
				Amount:       10,
				IsAnonymous:  true,
			},
			setup: func(donationMock *mocks.DonationStorer, fundMock *mocks.FundraiserStorer) {
				fundMock.On("GetFundraiserStatus", mock.Anything).Return(constants.ACTIVE_STATUS, nil).Once()
				donationMock.On("CreateDonation", mock.Anything).Return(uint(0), errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(donationRepo, fundRepo)

			// test service
			_, err := service.CreateDonation(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestListUserDonation(t *testing.T) {
	donationRepo := mocks.NewDonationStorer(t)
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(donationRepo, fundRepo)

	tests := []struct {
		name            string
		input           dto.ListUserDonationsRequest
		setup           func(donationMock *mocks.DonationStorer)
		isErrorExpected bool
	}{
		{
			name: "Success all donations",
			input: dto.ListUserDonationsRequest{
				UserId:             1,
				Search:             "donation",
				IsAnonymous:        "true",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListUserDonationsCount", mock.Anything).Return(uint(10), nil).Once()
				donationMock.On("ListUserDonations", mock.Anything).Return([]dto.DonationView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Success all donation with empty is_anonymous",
			input: dto.ListUserDonationsRequest{
				UserId:             1,
				Search:             "donation",
				IsAnonymous:        "",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListUserDonationsCount", mock.Anything).Return(uint(10), nil).Once()
				donationMock.On("ListUserDonations", mock.Anything).Return([]dto.DonationView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail user donations count",
			input: dto.ListUserDonationsRequest{
				UserId:             1,
				Search:             "donation",
				IsAnonymous:        "true",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListUserDonationsCount", mock.Anything).Return(uint(0), errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail user donations",
			input: dto.ListUserDonationsRequest{
				UserId:             1,
				Search:             "donation",
				IsAnonymous:        "",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListUserDonationsCount", mock.Anything).Return(uint(0), nil).Once()
				donationMock.On("ListUserDonations", mock.Anything).Return([]dto.DonationView{}, errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(donationRepo)

			// test service
			_, _, err := service.ListUserDonation(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestListFundraiserDonations(t *testing.T) {
	donationRepo := mocks.NewDonationStorer(t)
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(donationRepo, fundRepo)

	tests := []struct {
		name            string
		input           dto.ListFundraiserDonationsRequest
		setup           func(donationMock *mocks.DonationStorer)
		isErrorExpected bool
	}{
		{
			name: "Success all donations of fundraiser",
			input: dto.ListFundraiserDonationsRequest{
				FundraiserId: 1,
				Offset:       0,
				Limit:        100,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListFundraiserDonationsCount", mock.Anything).Return(uint(10), nil).Once()
				donationMock.On("ListFundraiserDonations", mock.Anything).Return([]dto.FundraiserDonationView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail fundraiser donations count",
			input: dto.ListFundraiserDonationsRequest{
				FundraiserId: 1,
				Offset:       0,
				Limit:        100,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListFundraiserDonationsCount", mock.Anything).Return(uint(0), errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail fundraiser donations",
			input: dto.ListFundraiserDonationsRequest{
				FundraiserId: 1,
				Offset:       0,
				Limit:        100,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListFundraiserDonationsCount", mock.Anything).Return(uint(0), nil).Once()
				donationMock.On("ListFundraiserDonations", mock.Anything).Return([]dto.FundraiserDonationView{}, errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(donationRepo)

			// test service
			_, _, err := service.ListFundraiserDonations(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestListDonations(t *testing.T) {
	donationRepo := mocks.NewDonationStorer(t)
	fundRepo := mocks.NewFundraiserStorer(t)
	service := NewService(donationRepo, fundRepo)

	tests := []struct {
		name            string
		input           dto.ListDonationsRequest
		setup           func(donationMock *mocks.DonationStorer)
		isErrorExpected bool
	}{
		{
			name: "Success all donations",
			input: dto.ListDonationsRequest{
				Search:             "donation",
				IsAnonymous:        "true",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListDonationsCount", mock.Anything).Return(uint(10), nil).Once()
				donationMock.On("ListDonations", mock.Anything).Return([]dto.FundraiserDonationView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Success all donation with empty is_anonymous",
			input: dto.ListDonationsRequest{
				Search:             "donation",
				IsAnonymous:        "",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListDonationsCount", mock.Anything).Return(uint(10), nil).Once()
				donationMock.On("ListDonations", mock.Anything).Return([]dto.FundraiserDonationView{}, nil).Once()
			},
			isErrorExpected: false,
		},
		{
			name: "Fail user donations count",
			input: dto.ListDonationsRequest{
				Search:             "donation",
				IsAnonymous:        "true",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListDonationsCount", mock.Anything).Return(uint(0), errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "Fail user donations",
			input: dto.ListDonationsRequest{
				Search:             "donation",
				IsAnonymous:        "",
				Offset:             0,
				Limit:              100,
				OrderByKey:         "first_name",
				OrderByIsAscending: true,
			},
			setup: func(donationMock *mocks.DonationStorer) {
				donationMock.On("GetListDonationsCount", mock.Anything).Return(uint(0), nil).Once()
				donationMock.On("ListDonations", mock.Anything).Return([]dto.FundraiserDonationView{}, errors.New("Error")).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(donationRepo)

			// test service
			_, _, err := service.ListDonations(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
