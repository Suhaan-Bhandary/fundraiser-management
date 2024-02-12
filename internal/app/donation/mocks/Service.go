// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	dto "github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateDonation provides a mock function with given fields: donationDetail
func (_m *Service) CreateDonation(donationDetail dto.CreateDonationRequest) (uint, error) {
	ret := _m.Called(donationDetail)

	if len(ret) == 0 {
		panic("no return value specified for CreateDonation")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.CreateDonationRequest) (uint, error)); ok {
		return rf(donationDetail)
	}
	if rf, ok := ret.Get(0).(func(dto.CreateDonationRequest) uint); ok {
		r0 = rf(donationDetail)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(dto.CreateDonationRequest) error); ok {
		r1 = rf(donationDetail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDonations provides a mock function with given fields: req
func (_m *Service) ListDonations(req dto.ListDonationsRequest) ([]dto.FundraiserDonationView, uint, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for ListDonations")
	}

	var r0 []dto.FundraiserDonationView
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(dto.ListDonationsRequest) ([]dto.FundraiserDonationView, uint, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(dto.ListDonationsRequest) []dto.FundraiserDonationView); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.FundraiserDonationView)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.ListDonationsRequest) uint); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(dto.ListDonationsRequest) error); ok {
		r2 = rf(req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListFundraiserDonations provides a mock function with given fields: req
func (_m *Service) ListFundraiserDonations(req dto.ListFundraiserDonationsRequest) ([]dto.FundraiserDonationView, uint, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for ListFundraiserDonations")
	}

	var r0 []dto.FundraiserDonationView
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(dto.ListFundraiserDonationsRequest) ([]dto.FundraiserDonationView, uint, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(dto.ListFundraiserDonationsRequest) []dto.FundraiserDonationView); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.FundraiserDonationView)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.ListFundraiserDonationsRequest) uint); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(dto.ListFundraiserDonationsRequest) error); ok {
		r2 = rf(req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListUserDonation provides a mock function with given fields: req
func (_m *Service) ListUserDonation(req dto.ListUserDonationsRequest) ([]dto.DonationView, uint, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for ListUserDonation")
	}

	var r0 []dto.DonationView
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(dto.ListUserDonationsRequest) ([]dto.DonationView, uint, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(dto.ListUserDonationsRequest) []dto.DonationView); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.DonationView)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.ListUserDonationsRequest) uint); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(dto.ListUserDonationsRequest) error); ok {
		r2 = rf(req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
