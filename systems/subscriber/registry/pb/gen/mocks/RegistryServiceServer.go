// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	gen "github.com/ukama/ukama/systems/subscriber/registry/pb/gen"
)

// RegistryServiceServer is an autogenerated mock type for the RegistryServiceServer type
type RegistryServiceServer struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *RegistryServiceServer) Add(_a0 context.Context, _a1 *gen.AddSubscriberRequest) (*gen.AddSubscriberResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.AddSubscriberResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.AddSubscriberRequest) (*gen.AddSubscriberResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.AddSubscriberRequest) *gen.AddSubscriberResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.AddSubscriberResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.AddSubscriberRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *RegistryServiceServer) Delete(_a0 context.Context, _a1 *gen.DeleteSubscriberRequest) (*gen.DeleteSubscriberResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.DeleteSubscriberResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.DeleteSubscriberRequest) (*gen.DeleteSubscriberResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.DeleteSubscriberRequest) *gen.DeleteSubscriberResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.DeleteSubscriberResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.DeleteSubscriberRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *RegistryServiceServer) Get(_a0 context.Context, _a1 *gen.GetSubscriberRequest) (*gen.GetSubscriberResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetSubscriberResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetSubscriberRequest) (*gen.GetSubscriberResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetSubscriberRequest) *gen.GetSubscriberResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetSubscriberResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetSubscriberRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByNetwork provides a mock function with given fields: _a0, _a1
func (_m *RegistryServiceServer) GetByNetwork(_a0 context.Context, _a1 *gen.GetByNetworkRequest) (*gen.GetByNetworkResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetByNetworkResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetByNetworkRequest) (*gen.GetByNetworkResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetByNetworkRequest) *gen.GetByNetworkResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetByNetworkResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetByNetworkRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSubscribers provides a mock function with given fields: _a0, _a1
func (_m *RegistryServiceServer) ListSubscribers(_a0 context.Context, _a1 *gen.ListSubscribersRequest) (*gen.ListSubscribersResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListSubscribersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListSubscribersRequest) (*gen.ListSubscribersResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListSubscribersRequest) *gen.ListSubscribersResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListSubscribersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListSubscribersRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *RegistryServiceServer) Update(_a0 context.Context, _a1 *gen.UpdateSubscriberRequest) (*gen.UpdateSubscriberResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpdateSubscriberResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpdateSubscriberRequest) (*gen.UpdateSubscriberResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpdateSubscriberRequest) *gen.UpdateSubscriberResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpdateSubscriberResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpdateSubscriberRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedRegistryServiceServer provides a mock function with given fields:
func (_m *RegistryServiceServer) mustEmbedUnimplementedRegistryServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewRegistryServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewRegistryServiceServer creates a new instance of RegistryServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRegistryServiceServer(t mockConstructorTestingTNewRegistryServiceServer) *RegistryServiceServer {
	mock := &RegistryServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
