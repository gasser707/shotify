// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/gasser707/go-gql-server/graphql/model"
	mock "github.com/stretchr/testify/mock"
)

// AuthServiceInterface is an autogenerated mock type for the AuthServiceInterface type
type AuthServiceInterface struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, input
func (_m *AuthServiceInterface) Login(ctx context.Context, input model.LoginInput) (bool, error) {
	ret := _m.Called(ctx, input)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, model.LoginInput) bool); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.LoginInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx
func (_m *AuthServiceInterface) Logout(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: ctx
func (_m *AuthServiceInterface) Refresh(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateCredentials provides a mock function with given fields: c
func (_m *AuthServiceInterface) ValidateCredentials(c context.Context) (int, model.Role, error) {
	ret := _m.Called(c)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 model.Role
	if rf, ok := ret.Get(1).(func(context.Context) model.Role); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Get(1).(model.Role)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(c)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}