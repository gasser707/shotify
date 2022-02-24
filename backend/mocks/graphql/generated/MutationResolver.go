// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	custom "github.com/gasser707/go-gql-server/graphql/custom"

	mock "github.com/stretchr/testify/mock"

	model "github.com/gasser707/go-gql-server/graphql/model"
)

// MutationResolver is an autogenerated mock type for the MutationResolver type
type MutationResolver struct {
	mock.Mock
}

// AutoGenerateLabels provides a mock function with given fields: ctx, id
func (_m *MutationResolver) AutoGenerateLabels(ctx context.Context, id string) ([]string, error) {
	ret := _m.Called(ctx, id)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuyImage provides a mock function with given fields: ctx, id
func (_m *MutationResolver) BuyImage(ctx context.Context, id string) (*custom.Sale, error) {
	ret := _m.Called(ctx, id)

	var r0 *custom.Sale
	if rf, ok := ret.Get(0).(func(context.Context, string) *custom.Sale); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*custom.Sale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteImages provides a mock function with given fields: ctx, input
func (_m *MutationResolver) DeleteImages(ctx context.Context, input []string) (bool, error) {
	ret := _m.Called(ctx, input)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, []string) bool); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, input
func (_m *MutationResolver) Login(ctx context.Context, input model.LoginInput) (bool, error) {
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

// Logout provides a mock function with given fields: ctx, input
func (_m *MutationResolver) Logout(ctx context.Context, input *bool) (bool, error) {
	ret := _m.Called(ctx, input)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *bool) bool); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *bool) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogoutAll provides a mock function with given fields: ctx, input
func (_m *MutationResolver) LogoutAll(ctx context.Context, input *bool) (bool, error) {
	ret := _m.Called(ctx, input)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *bool) bool); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *bool) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessPasswordReset provides a mock function with given fields: ctx, resetToken, newPassword
func (_m *MutationResolver) ProcessPasswordReset(ctx context.Context, resetToken string, newPassword string) (bool, error) {
	ret := _m.Called(ctx, resetToken, newPassword)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, resetToken, newPassword)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, resetToken, newPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: ctx, input
func (_m *MutationResolver) Refresh(ctx context.Context, input *bool) (bool, error) {
	ret := _m.Called(ctx, input)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *bool) bool); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *bool) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, input
func (_m *MutationResolver) RegisterUser(ctx context.Context, input model.NewUserInput) (*custom.User, error) {
	ret := _m.Called(ctx, input)

	var r0 *custom.User
	if rf, ok := ret.Get(0).(func(context.Context, model.NewUserInput) *custom.User); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*custom.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.NewUserInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestPasswordReset provides a mock function with given fields: ctx, email
func (_m *MutationResolver) RequestPasswordReset(ctx context.Context, email string) (bool, error) {
	ret := _m.Called(ctx, email)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateImage provides a mock function with given fields: ctx, input
func (_m *MutationResolver) UpdateImage(ctx context.Context, input model.UpdateImageInput) (*custom.Image, error) {
	ret := _m.Called(ctx, input)

	var r0 *custom.Image
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateImageInput) *custom.Image); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*custom.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UpdateImageInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, input
func (_m *MutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*custom.User, error) {
	ret := _m.Called(ctx, input)

	var r0 *custom.User
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateUserInput) *custom.User); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*custom.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UpdateUserInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadImages provides a mock function with given fields: ctx, input
func (_m *MutationResolver) UploadImages(ctx context.Context, input []*model.NewImageInput) ([]*custom.Image, error) {
	ret := _m.Called(ctx, input)

	var r0 []*custom.Image
	if rf, ok := ret.Get(0).(func(context.Context, []*model.NewImageInput) []*custom.Image); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*custom.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*model.NewImageInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateUser provides a mock function with given fields: ctx, validationToken
func (_m *MutationResolver) ValidateUser(ctx context.Context, validationToken string) (bool, error) {
	ret := _m.Called(ctx, validationToken)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, validationToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, validationToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}