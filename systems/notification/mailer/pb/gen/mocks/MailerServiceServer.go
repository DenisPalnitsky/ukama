// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	gen "github.com/ukama/ukama/systems/notification/mailer/pb/gen"
)

// MailerServiceServer is an autogenerated mock type for the MailerServiceServer type
type MailerServiceServer struct {
	mock.Mock
}

// GetEmailById provides a mock function with given fields: _a0, _a1
func (_m *MailerServiceServer) GetEmailById(_a0 context.Context, _a1 *gen.GetEmailByIdRequest) (*gen.GetEmailByIdResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetEmailByIdResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetEmailByIdRequest) (*gen.GetEmailByIdResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetEmailByIdRequest) *gen.GetEmailByIdResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetEmailByIdResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetEmailByIdRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendEmail provides a mock function with given fields: _a0, _a1
func (_m *MailerServiceServer) SendEmail(_a0 context.Context, _a1 *gen.SendEmailRequest) (*gen.SendEmailResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.SendEmailResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SendEmailRequest) (*gen.SendEmailResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SendEmailRequest) *gen.SendEmailResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.SendEmailResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.SendEmailRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedMailerServiceServer provides a mock function with given fields:
func (_m *MailerServiceServer) mustEmbedUnimplementedMailerServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewMailerServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMailerServiceServer creates a new instance of MailerServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMailerServiceServer(t mockConstructorTestingTNewMailerServiceServer) *MailerServiceServer {
	mock := &MailerServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
