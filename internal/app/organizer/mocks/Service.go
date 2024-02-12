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

// DeleteOrganizer provides a mock function with given fields: organizerId
func (_m *Service) DeleteOrganizer(organizerId uint) error {
	ret := _m.Called(organizerId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOrganizer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(organizerId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOrganizer provides a mock function with given fields: organizerId
func (_m *Service) GetOrganizer(organizerId uint) (dto.OrganizerView, error) {
	ret := _m.Called(organizerId)

	if len(ret) == 0 {
		panic("no return value specified for GetOrganizer")
	}

	var r0 dto.OrganizerView
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (dto.OrganizerView, error)); ok {
		return rf(organizerId)
	}
	if rf, ok := ret.Get(0).(func(uint) dto.OrganizerView); ok {
		r0 = rf(organizerId)
	} else {
		r0 = ret.Get(0).(dto.OrganizerView)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(organizerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizerList provides a mock function with given fields: req
func (_m *Service) GetOrganizerList(req dto.ListOrganizersRequest) ([]dto.OrganizerView, uint, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for GetOrganizerList")
	}

	var r0 []dto.OrganizerView
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(dto.ListOrganizersRequest) ([]dto.OrganizerView, uint, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(dto.ListOrganizersRequest) []dto.OrganizerView); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.OrganizerView)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.ListOrganizersRequest) uint); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(dto.ListOrganizersRequest) error); ok {
		r2 = rf(req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LoginOrganizer provides a mock function with given fields: req
func (_m *Service) LoginOrganizer(req dto.LoginOrganizerRequest) (string, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for LoginOrganizer")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.LoginOrganizerRequest) (string, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(dto.LoginOrganizerRequest) string); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(dto.LoginOrganizerRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterOrganizer provides a mock function with given fields: userDetail
func (_m *Service) RegisterOrganizer(userDetail dto.RegisterOrganizerRequest) error {
	ret := _m.Called(userDetail)

	if len(ret) == 0 {
		panic("no return value specified for RegisterOrganizer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.RegisterOrganizerRequest) error); ok {
		r0 = rf(userDetail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrganizer provides a mock function with given fields: req
func (_m *Service) UpdateOrganizer(req dto.UpdateOrganizerRequest) error {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrganizer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.UpdateOrganizerRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyOrganizer provides a mock function with given fields: organizerId
func (_m *Service) VerifyOrganizer(organizerId uint) error {
	ret := _m.Called(organizerId)

	if len(ret) == 0 {
		panic("no return value specified for VerifyOrganizer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(organizerId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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