// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	dto "github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// OrganizerStorer is an autogenerated mock type for the OrganizerStorer type
type OrganizerStorer struct {
	mock.Mock
}

// DeleteOrganizer provides a mock function with given fields: organizerId
func (_m *OrganizerStorer) DeleteOrganizer(organizerId uint) error {
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
func (_m *OrganizerStorer) GetOrganizer(organizerId uint) (dto.OrganizerView, error) {
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

// GetOrganizerIDPassword provides a mock function with given fields: email
func (_m *OrganizerStorer) GetOrganizerIDPassword(email string) (uint, string, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetOrganizerIDPassword")
	}

	var r0 uint
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (uint, string, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) uint); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(email)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetOrganizerList provides a mock function with given fields: search, verified
func (_m *OrganizerStorer) GetOrganizerList(search string, verified string) ([]dto.OrganizerView, error) {
	ret := _m.Called(search, verified)

	if len(ret) == 0 {
		panic("no return value specified for GetOrganizerList")
	}

	var r0 []dto.OrganizerView
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]dto.OrganizerView, error)); ok {
		return rf(search, verified)
	}
	if rf, ok := ret.Get(0).(func(string, string) []dto.OrganizerView); ok {
		r0 = rf(search, verified)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.OrganizerView)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(search, verified)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterOrganizer provides a mock function with given fields: orgDetail
func (_m *OrganizerStorer) RegisterOrganizer(orgDetail dto.RegisterOrganizerRequest) error {
	ret := _m.Called(orgDetail)

	if len(ret) == 0 {
		panic("no return value specified for RegisterOrganizer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.RegisterOrganizerRequest) error); ok {
		r0 = rf(orgDetail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrganizer provides a mock function with given fields: req
func (_m *OrganizerStorer) UpdateOrganizer(req dto.UpdateOrganizerRequest) error {
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
func (_m *OrganizerStorer) VerifyOrganizer(organizerId uint) error {
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

// NewOrganizerStorer creates a new instance of OrganizerStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrganizerStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrganizerStorer {
	mock := &OrganizerStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
