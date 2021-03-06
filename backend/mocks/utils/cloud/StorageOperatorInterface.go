// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// StorageOperatorInterface is an autogenerated mock type for the StorageOperatorInterface type
type StorageOperatorInterface struct {
	mock.Mock
}

// DeleteImage provides a mock function with given fields: path
func (_m *StorageOperatorInterface) DeleteImage(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadImage provides a mock function with given fields: img, imgName, productId
func (_m *StorageOperatorInterface) UploadImage(img io.Reader, imgName string, productId string) (string, error) {
	ret := _m.Called(img, imgName, productId)

	var r0 string
	if rf, ok := ret.Get(0).(func(io.Reader, string, string) string); ok {
		r0 = rf(img, imgName, productId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader, string, string) error); ok {
		r1 = rf(img, imgName, productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
