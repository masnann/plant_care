// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/masnann/plant_care/features/user/domain"
	mock "github.com/stretchr/testify/mock"
)

// RepoUserInterface is an autogenerated mock type for the RepoUserInterface type
type RepoUserInterface struct {
	mock.Mock
}

// GetAllUsers provides a mock function with given fields:
func (_m *RepoUserInterface) GetAllUsers() ([]*domain.UserModel, error) {
	ret := _m.Called()

	var r0 []*domain.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*domain.UserModel, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*domain.UserModel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *RepoUserInterface) GetUserByEmail(email string) (*domain.UserModel, error) {
	ret := _m.Called(email)

	var r0 *domain.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.UserModel, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.UserModel); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserById provides a mock function with given fields: userId
func (_m *RepoUserInterface) GetUserById(userId uint64) (*domain.UserModel, error) {
	ret := _m.Called(userId)

	var r0 *domain.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*domain.UserModel, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uint64) *domain.UserModel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepoUserInterface creates a new instance of RepoUserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepoUserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepoUserInterface {
	mock := &RepoUserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
