// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/masnann/plant_care/features/user/domain"
	mock "github.com/stretchr/testify/mock"
)

// ServiceAuthInterface is an autogenerated mock type for the ServiceAuthInterface type
type ServiceAuthInterface struct {
	mock.Mock
}

// Login provides a mock function with given fields: email, password
func (_m *ServiceAuthInterface) Login(email string, password string) (*domain.UserModel, string, error) {
	ret := _m.Called(email, password)

	var r0 *domain.UserModel
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (*domain.UserModel, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *domain.UserModel); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: newData
func (_m *ServiceAuthInterface) Register(newData *domain.UserModel) (*domain.UserModel, error) {
	ret := _m.Called(newData)

	var r0 *domain.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.UserModel) (*domain.UserModel, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(*domain.UserModel) *domain.UserModel); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.UserModel) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServiceAuthInterface creates a new instance of ServiceAuthInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceAuthInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceAuthInterface {
	mock := &ServiceAuthInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
