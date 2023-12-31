// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// HandlerNotificationInterface is an autogenerated mock type for the HandlerNotificationInterface type
type HandlerNotificationInterface struct {
	mock.Mock
}

// GetPaginationNotifications provides a mock function with given fields:
func (_m *HandlerNotificationInterface) GetPaginationNotifications() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// NewHandlerNotificationInterface creates a new instance of HandlerNotificationInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandlerNotificationInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *HandlerNotificationInterface {
	mock := &HandlerNotificationInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
