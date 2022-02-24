// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	custom "github.com/gasser707/go-gql-server/graphql/custom"
	mock "github.com/stretchr/testify/mock"

	model "github.com/gasser707/go-gql-server/graphql/model"
)

// ImagesServiceInterface is an autogenerated mock type for the ImagesServiceInterface type
type ImagesServiceInterface struct {
	mock.Mock
}

// AutoGenerateLabels provides a mock function with given fields: ctx, imageId
func (_m *ImagesServiceInterface) AutoGenerateLabels(ctx context.Context, imageId string) ([]string, error) {
	ret := _m.Called(ctx, imageId)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, imageId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, imageId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteImages provides a mock function with given fields: ctx, input
func (_m *ImagesServiceInterface) DeleteImages(ctx context.Context, input []string) (bool, error) {
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

// GetImageById provides a mock function with given fields: ctx, ID
func (_m *ImagesServiceInterface) GetImageById(ctx context.Context, ID string) (*custom.Image, error) {
	ret := _m.Called(ctx, ID)

	var r0 *custom.Image
	if rf, ok := ret.Get(0).(func(context.Context, string) *custom.Image); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*custom.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImages provides a mock function with given fields: ctx, input
func (_m *ImagesServiceInterface) GetImages(ctx context.Context, input *model.ImageFilterInput) ([]*custom.Image, error) {
	ret := _m.Called(ctx, input)

	var r0 []*custom.Image
	if rf, ok := ret.Get(0).(func(context.Context, *model.ImageFilterInput) []*custom.Image); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*custom.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.ImageFilterInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateImage provides a mock function with given fields: ctx, input
func (_m *ImagesServiceInterface) UpdateImage(ctx context.Context, input *model.UpdateImageInput) (*custom.Image, error) {
	ret := _m.Called(ctx, input)

	var r0 *custom.Image
	if rf, ok := ret.Get(0).(func(context.Context, *model.UpdateImageInput) *custom.Image); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*custom.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.UpdateImageInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadImages provides a mock function with given fields: ctx, input
func (_m *ImagesServiceInterface) UploadImages(ctx context.Context, input []*model.NewImageInput) ([]*custom.Image, error) {
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
