// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/superhero-match/consumer-delete-media/internal/db/model"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// DeleteProfilePicture provides a mock function with given fields: pp
func (_m *DB) DeleteProfilePicture(pp model.ProfilePicture) error {
	ret := _m.Called(pp)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.ProfilePicture) error); ok {
		r0 = rf(pp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
