// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	models "github.com/MorZLE/url-shortener/internal/models"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// CheckPing provides a mock function with given fields: c
func (_m *Handler) CheckPing(c *gin.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gin.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cookie provides a mock function with given fields: c
func (_m *Handler) Cookie(c *gin.Context) string {
	ret := _m.Called(c)

	var r0 string
	if rf, ok := ret.Get(0).(func(*gin.Context) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// JSONURLShort provides a mock function with given fields: c, obj
func (_m *Handler) JSONURLShort(c *gin.Context, obj models.URLShort) {
	_m.Called(c, obj)
}

// JSONURLShortBatch provides a mock function with given fields: c
func (_m *Handler) JSONURLShortBatch(c *gin.Context) {
	_m.Called(c)
}

// RunServer provides a mock function with given fields:
func (_m *Handler) RunServer() {
	_m.Called()
}

// URLGetCookie provides a mock function with given fields: c
func (_m *Handler) URLGetCookie(c *gin.Context) {
	_m.Called(c)
}

// URLGetID provides a mock function with given fields: c
func (_m *Handler) URLGetID(c *gin.Context) {
	_m.Called(c)
}

// URLShortener provides a mock function with given fields: c
func (_m *Handler) URLShortener(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandler(t mockConstructorTestingTNewHandler) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
