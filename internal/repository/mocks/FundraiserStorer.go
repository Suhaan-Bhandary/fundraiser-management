// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FundraiserStorer is an autogenerated mock type for the FundraiserStorer type
type FundraiserStorer struct {
	mock.Mock
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
