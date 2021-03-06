// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/gasser707/go-gql-server/utils/auth"

	mock "github.com/stretchr/testify/mock"

	model "github.com/gasser707/go-gql-server/graphql/model"
)

// TokenOperatorInterface is an autogenerated mock type for the TokenOperatorInterface type
type TokenOperatorInterface struct {
	mock.Mock
}

// CreateStatelessToken provides a mock function with given fields: userId, kind
func (_m *TokenOperatorInterface) CreateStatelessToken(userId string, kind auth.StatelessToken) (string, error) {
	ret := _m.Called(userId, kind)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, auth.StatelessToken) string); ok {
		r0 = rf(userId, kind)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, auth.StatelessToken) error); ok {
		r1 = rf(userId, kind)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTokens provides a mock function with given fields: userId, userRole
func (_m *TokenOperatorInterface) CreateTokens(userId string, userRole model.Role) (*auth.TokenDetails, error) {
	ret := _m.Called(userId, userRole)

	var r0 *auth.TokenDetails
	if rf, ok := ret.Get(0).(func(string, model.Role) *auth.TokenDetails); ok {
		r0 = rf(userId, userRole)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.TokenDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.Role) error); ok {
		r1 = rf(userId, userRole)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtractAccessTokenMetadata provides a mock function with given fields: c
func (_m *TokenOperatorInterface) ExtractAccessTokenMetadata(c context.Context) (*auth.AccessDetails, error) {
	ret := _m.Called(c)

	var r0 *auth.AccessDetails
	if rf, ok := ret.Get(0).(func(context.Context) *auth.AccessDetails); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.AccessDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtractRefreshMetadata provides a mock function with given fields: ctx
func (_m *TokenOperatorInterface) ExtractRefreshMetadata(ctx context.Context) (*auth.RefreshDetails, error) {
	ret := _m.Called(ctx)

	var r0 *auth.RefreshDetails
	if rf, ok := ret.Get(0).(func(context.Context) *auth.RefreshDetails); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.RefreshDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtractStatelessTokenMetadata provides a mock function with given fields: ctx, tokenString, kind
func (_m *TokenOperatorInterface) ExtractStatelessTokenMetadata(ctx context.Context, tokenString string, kind auth.StatelessToken) (string, error) {
	ret := _m.Called(ctx, tokenString, kind)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, auth.StatelessToken) string); ok {
		r0 = rf(ctx, tokenString, kind)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, auth.StatelessToken) error); ok {
		r1 = rf(ctx, tokenString, kind)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
