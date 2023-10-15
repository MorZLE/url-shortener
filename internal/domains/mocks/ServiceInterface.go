// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/MorZLE/url-shortener/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// CheckPing provides a mock function with given fields:
func (_m *ServiceInterface) CheckPing() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// URLGetID provides a mock function with given fields: url
func (_m *ServiceInterface) URLGetID(url string) (string, error) {
	ret := _m.Called(url)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URLShorter provides a mock function with given fields: url
func (_m *ServiceInterface) URLShorter(url string) (string, error) {
	ret := _m.Called(url)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URLsShorter provides a mock function with given fields: url
func (_m *ServiceInterface) URLsShorter(url []models.BatchSet) ([]models.BatchGet, error) {
	ret := _m.Called(url)

	var r0 []models.BatchGet
	var r1 error
	if rf, ok := ret.Get(0).(func([]models.BatchSet) ([]models.BatchGet, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func([]models.BatchSet) []models.BatchGet); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.BatchGet)
		}
	}

	if rf, ok := ret.Get(1).(func([]models.BatchSet) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceInterface creates a new instance of ServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceInterface(t mockConstructorTestingTNewServiceInterface) *ServiceInterface {
	mock := &ServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
