// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/MorZLE/url-shortener/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CheckPing provides a mock function with given fields:
func (_m *Service) CheckPing() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Generate provides a mock function with given fields: num
func (_m *Service) Generate(num int) (string, error) {
	ret := _m.Called(num)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (string, error)); ok {
		return rf(num)
	}
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(num)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateCookie provides a mock function with given fields:
func (_m *Service) GenerateCookie() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetAllURLUsers provides a mock function with given fields: id
func (_m *Service) GetAllURLUsers(id string) ([]models.AllURLs, error) {
	ret := _m.Called(id)

	var r0 []models.AllURLs
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.AllURLs, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) []models.AllURLs); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.AllURLs)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URLGetID provides a mock function with given fields: url
func (_m *Service) URLGetID(url string) (string, error) {
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

// URLShorter provides a mock function with given fields: id, url
func (_m *Service) URLShorter(id string, url string) (string, error) {
	ret := _m.Called(id, url)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(id, url)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(id, url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URLsShorter provides a mock function with given fields: id, url
func (_m *Service) URLsShorter(id string, url []models.BatchSet) ([]models.BatchGet, error) {
	ret := _m.Called(id, url)

	var r0 []models.BatchGet
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []models.BatchSet) ([]models.BatchGet, error)); ok {
		return rf(id, url)
	}
	if rf, ok := ret.Get(0).(func(string, []models.BatchSet) []models.BatchGet); ok {
		r0 = rf(id, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.BatchGet)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []models.BatchSet) error); ok {
		r1 = rf(id, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
