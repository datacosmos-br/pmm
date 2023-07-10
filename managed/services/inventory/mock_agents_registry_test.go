// Code generated by mockery v2.31.1. DO NOT EDIT.

package inventory

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockAgentsRegistry is an autogenerated mock type for the agentsRegistry type
type mockAgentsRegistry struct {
	mock.Mock
}

// IsConnected provides a mock function with given fields: pmmAgentID
func (_m *mockAgentsRegistry) IsConnected(pmmAgentID string) bool {
	ret := _m.Called(pmmAgentID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(pmmAgentID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Kick provides a mock function with given fields: ctx, pmmAgentID
func (_m *mockAgentsRegistry) Kick(ctx context.Context, pmmAgentID string) {
	_m.Called(ctx, pmmAgentID)
}

// newMockAgentsRegistry creates a new instance of mockAgentsRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAgentsRegistry(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockAgentsRegistry {
	mock := &mockAgentsRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
