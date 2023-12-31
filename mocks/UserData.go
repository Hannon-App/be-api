// Code generated by mockery v2.35.3. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	users "Hannon-app/features/users"
)

// UserData is an autogenerated mock type for the UserDataInterface type
type UserData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: adminID, id
func (_m *UserData) Delete(adminID uint, id uint) error {
	ret := _m.Called(adminID, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(adminID, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: input, fileImages, fileID, filenameImages, filenameID
func (_m *UserData) Insert(input users.UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error {
	ret := _m.Called(input, fileImages, fileID, filenameImages, filenameID)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.UserCore, multipart.File, multipart.File, string, string) error); ok {
		r0 = rf(input, fileImages, fileID, filenameImages, filenameID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: email, password
func (_m *UserData) Login(email string, password string) (users.UserCore, error) {
	ret := _m.Called(email, password)

	var r0 users.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (users.UserCore, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) users.UserCore); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields: adminID, page, userPerPage, searchName
func (_m *UserData) ReadAll(adminID uint, page uint, userPerPage uint, searchName string) ([]users.UserCore, int64, error) {
	ret := _m.Called(adminID, page, userPerPage, searchName)

	var r0 []users.UserCore
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(uint, uint, uint, string) ([]users.UserCore, int64, error)); ok {
		return rf(adminID, page, userPerPage, searchName)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, uint, string) []users.UserCore); ok {
		r0 = rf(adminID, page, userPerPage, searchName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.UserCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint, uint, string) int64); ok {
		r1 = rf(adminID, page, userPerPage, searchName)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(uint, uint, uint, string) error); ok {
		r2 = rf(adminID, page, userPerPage, searchName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectById provides a mock function with given fields: id
func (_m *UserData) SelectById(id uint) (users.UserCore, error) {
	ret := _m.Called(id)

	var r0 users.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (users.UserCore, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) users.UserCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: uID, id, input, fileImages, fileID, filenameImages, filenameID
func (_m *UserData) UpdateUser(uID uint, id uint, input users.UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error {
	ret := _m.Called(uID, id, input, fileImages, fileID, filenameImages, filenameID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, users.UserCore, multipart.File, multipart.File, string, string) error); ok {
		r0 = rf(uID, id, input, fileImages, fileID, filenameImages, filenameID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserData(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
