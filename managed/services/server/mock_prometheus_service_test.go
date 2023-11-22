// Code generated by mockery. DO NOT EDIT.

package server

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockPrometheusService is an autogenerated mock type for the prometheusService type
type mockPrometheusService struct {
	mock.Mock
}

// IsReady provides a mock function with given fields: ctx
func (_m *mockPrometheusService) IsReady(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsReady")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestConfigurationUpdate provides a mock function with given fields:
func (_m *mockPrometheusService) RequestConfigurationUpdate() {
	_m.Called()
}

// newMockPrometheusService creates a new instance of mockPrometheusService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockPrometheusService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockPrometheusService {
	mock := &mockPrometheusService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
