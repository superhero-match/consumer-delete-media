// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// DeleteSuperhero provides a mock function with given fields: keys
func (_m *Cache) DeleteSuperhero(keys []string) error {
	ret := _m.Called(keys)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(keys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}