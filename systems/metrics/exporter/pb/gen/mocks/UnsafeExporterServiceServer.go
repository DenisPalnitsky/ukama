// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeExporterServiceServer is an autogenerated mock type for the UnsafeExporterServiceServer type
type UnsafeExporterServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedExporterServiceServer provides a mock function with given fields:
func (_m *UnsafeExporterServiceServer) mustEmbedUnimplementedExporterServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafeExporterServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeExporterServiceServer creates a new instance of UnsafeExporterServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeExporterServiceServer(t mockConstructorTestingTNewUnsafeExporterServiceServer) *UnsafeExporterServiceServer {
	mock := &UnsafeExporterServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
