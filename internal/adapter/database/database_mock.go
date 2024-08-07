// Code generated by mockery v2.42.3. DO NOT EDIT.

package database

import (
	health "github.com/jfelipearaujo-org/ms-customer-management/internal/shared/health"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// MockDatabaseService is an autogenerated mock type for the DatabaseService type
type MockDatabaseService struct {
	mock.Mock
}

// GetInstance provides a mock function with given fields:
func (_m *MockDatabaseService) GetInstance() *sql.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetInstance")
	}

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// Health provides a mock function with given fields:
func (_m *MockDatabaseService) Health() *health.HealthStatus {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Health")
	}

	var r0 *health.HealthStatus
	if rf, ok := ret.Get(0).(func() *health.HealthStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*health.HealthStatus)
		}
	}

	return r0
}

// NewMockDatabaseService creates a new instance of MockDatabaseService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDatabaseService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDatabaseService {
	mock := &MockDatabaseService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
