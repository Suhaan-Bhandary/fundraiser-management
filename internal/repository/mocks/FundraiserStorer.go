// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	dto "github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// FundraiserStorer is an autogenerated mock type for the FundraiserStorer type
type FundraiserStorer struct {
	mock.Mock
}

// BanFundraiser provides a mock function with given fields: fundraiserId
func (_m *FundraiserStorer) BanFundraiser(fundraiserId uint) error {
	ret := _m.Called(fundraiserId)

	if len(ret) == 0 {
		panic("no return value specified for BanFundraiser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(fundraiserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseFundraiser provides a mock function with given fields: fundraiserId
func (_m *FundraiserStorer) CloseFundraiser(fundraiserId uint) error {
	ret := _m.Called(fundraiserId)

	if len(ret) == 0 {
		panic("no return value specified for CloseFundraiser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(fundraiserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateFundraiser provides a mock function with given fields: fundDetail
func (_m *FundraiserStorer) CreateFundraiser(fundDetail dto.CreateFundraiserRequest) (uint, error) {
	ret := _m.Called(fundDetail)

	if len(ret) == 0 {
		panic("no return value specified for CreateFundraiser")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.CreateFundraiserRequest) (uint, error)); ok {
		return rf(fundDetail)
	}
	if rf, ok := ret.Get(0).(func(dto.CreateFundraiserRequest) uint); ok {
		r0 = rf(fundDetail)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(dto.CreateFundraiserRequest) error); ok {
		r1 = rf(fundDetail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFundraiser provides a mock function with given fields: fundraiserId
func (_m *FundraiserStorer) DeleteFundraiser(fundraiserId uint) error {
	ret := _m.Called(fundraiserId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFundraiser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(fundraiserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFundraiser provides a mock function with given fields: fundraiserId
func (_m *FundraiserStorer) GetFundraiser(fundraiserId uint) (dto.FundraiserView, error) {
	ret := _m.Called(fundraiserId)

	if len(ret) == 0 {
		panic("no return value specified for GetFundraiser")
	}

	var r0 dto.FundraiserView
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (dto.FundraiserView, error)); ok {
		return rf(fundraiserId)
	}
	if rf, ok := ret.Get(0).(func(uint) dto.FundraiserView); ok {
		r0 = rf(fundraiserId)
	} else {
		r0 = ret.Get(0).(dto.FundraiserView)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(fundraiserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFundraiserOrganizerId provides a mock function with given fields: fundraiserId
func (_m *FundraiserStorer) GetFundraiserOrganizerId(fundraiserId uint) (uint, error) {
	ret := _m.Called(fundraiserId)

	if len(ret) == 0 {
		panic("no return value specified for GetFundraiserOrganizerId")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (uint, error)); ok {
		return rf(fundraiserId)
	}
	if rf, ok := ret.Get(0).(func(uint) uint); ok {
		r0 = rf(fundraiserId)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(fundraiserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFundraiserOrganizerIdAndStatus provides a mock function with given fields: fundraiserId
func (_m *FundraiserStorer) GetFundraiserOrganizerIdAndStatus(fundraiserId uint) (uint, string, error) {
	ret := _m.Called(fundraiserId)

	if len(ret) == 0 {
		panic("no return value specified for GetFundraiserOrganizerIdAndStatus")
	}

	var r0 uint
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(uint) (uint, string, error)); ok {
		return rf(fundraiserId)
	}
	if rf, ok := ret.Get(0).(func(uint) uint); ok {
		r0 = rf(fundraiserId)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(uint) string); ok {
		r1 = rf(fundraiserId)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(uint) error); ok {
		r2 = rf(fundraiserId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetFundraiserStatus provides a mock function with given fields: fundraiserId
func (_m *FundraiserStorer) GetFundraiserStatus(fundraiserId uint) (string, error) {
	ret := _m.Called(fundraiserId)

	if len(ret) == 0 {
		panic("no return value specified for GetFundraiserStatus")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (string, error)); ok {
		return rf(fundraiserId)
	}
	if rf, ok := ret.Get(0).(func(uint) string); ok {
		r0 = rf(fundraiserId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(fundraiserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListFundraisersCount provides a mock function with given fields: req
func (_m *FundraiserStorer) GetListFundraisersCount(req dto.ListFundraisersRequest) (uint, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for GetListFundraisersCount")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.ListFundraisersRequest) (uint, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(dto.ListFundraisersRequest) uint); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(dto.ListFundraisersRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFundraiser provides a mock function with given fields: req
func (_m *FundraiserStorer) ListFundraiser(req dto.ListFundraisersRequest) ([]dto.FundraiserView, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for ListFundraiser")
	}

	var r0 []dto.FundraiserView
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.ListFundraisersRequest) ([]dto.FundraiserView, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(dto.ListFundraisersRequest) []dto.FundraiserView); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.FundraiserView)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.ListFundraisersRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnBanFundraiser provides a mock function with given fields: fundraiserId
func (_m *FundraiserStorer) UnBanFundraiser(fundraiserId uint) error {
	ret := _m.Called(fundraiserId)

	if len(ret) == 0 {
		panic("no return value specified for UnBanFundraiser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(fundraiserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFundraiser provides a mock function with given fields: updateDetail
func (_m *FundraiserStorer) UpdateFundraiser(updateDetail dto.UpdateFundraiserRequest) error {
	ret := _m.Called(updateDetail)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFundraiser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.UpdateFundraiserRequest) error); ok {
		r0 = rf(updateDetail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFundraiserStorer creates a new instance of FundraiserStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFundraiserStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *FundraiserStorer {
	mock := &FundraiserStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
