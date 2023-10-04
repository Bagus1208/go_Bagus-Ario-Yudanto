package mock

import (
	"clean_architecture/features/users/entity"

	mock "github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (_m *RepositoryMock) Insert(data entity.User) (entity.User, error) {
	ret := _m.Called(data)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RepositoryMock) GetAll() ([]entity.User, error) {
	ret := _m.Called()

	var r0 []entity.User
	if rf, ok := ret.Get(0).(func() []entity.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r0 = ret.Get(1).([]entity.User)
		}
	}

	return r0, r1
}

func (_m *RepositoryMock) Login(email string, password string) *entity.User {
	ret := _m.Called(email, password)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string, string) *entity.User); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(*entity.User)
	}

	return r0
}
