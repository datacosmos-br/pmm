// Code generated by mockery v2.31.1. DO NOT EDIT.

package dbaas

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockGrafanaClient is an autogenerated mock type for the grafanaClient type
type mockGrafanaClient struct {
	mock.Mock
}

// CreateAdminAPIKey provides a mock function with given fields: ctx, name
func (_m *mockGrafanaClient) CreateAdminAPIKey(ctx context.Context, name string) (int64, string, error) {
	ret := _m.Called(ctx, name)

	var r0 int64
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, string, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) string); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, name)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteAPIKeyByID provides a mock function with given fields: ctx, id
func (_m *mockGrafanaClient) DeleteAPIKeyByID(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAPIKeysWithPrefix provides a mock function with given fields: ctx, name
func (_m *mockGrafanaClient) DeleteAPIKeysWithPrefix(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockGrafanaClient creates a new instance of mockGrafanaClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockGrafanaClient(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockGrafanaClient {
	mock := &mockGrafanaClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
