// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	emails "github.com/gasser707/go-gql-server/utils/emails"
	mock "github.com/stretchr/testify/mock"
)

// ReceiptEmailInterface is an autogenerated mock type for the ReceiptEmailInterface type
type ReceiptEmailInterface struct {
	mock.Mock
}

// GetImageID provides a mock function with given fields:
func (_m *ReceiptEmailInterface) GetImageID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetImageTitle provides a mock function with given fields:
func (_m *ReceiptEmailInterface) GetImageTitle() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *ReceiptEmailInterface) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPaymentMethod provides a mock function with given fields:
func (_m *ReceiptEmailInterface) GetPaymentMethod() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetSellerName provides a mock function with given fields:
func (_m *ReceiptEmailInterface) GetSellerName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetSender provides a mock function with given fields:
func (_m *ReceiptEmailInterface) GetSender() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTo provides a mock function with given fields:
func (_m *ReceiptEmailInterface) GetTo() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetType provides a mock function with given fields:
func (_m *ReceiptEmailInterface) GetType() emails.EmailType {
	ret := _m.Called()

	var r0 emails.EmailType
	if rf, ok := ret.Get(0).(func() emails.EmailType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(emails.EmailType)
	}

	return r0
}