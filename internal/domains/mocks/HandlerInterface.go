// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HandlerInterface is an autogenerated mock type for the HandlerInterface type
type HandlerInterface struct {
	mock.Mock
}

// JSONURLShort provides a mock function with given fields: w, r
func (_m *HandlerInterface) JSONURLShort(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// RunServer provides a mock function with given fields:
func (_m *HandlerInterface) RunServer() {
	_m.Called()
}

// URLGetID provides a mock function with given fields: w, r
func (_m *HandlerInterface) URLGetID(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// URLShortener provides a mock function with given fields: w, r
func (_m *HandlerInterface) URLShortener(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

type mockConstructorTestingTNewHandlerInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandlerInterface creates a new instance of HandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandlerInterface(t mockConstructorTestingTNewHandlerInterface) *HandlerInterface {
	mock := &HandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}