package mocks

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (_m *MockUserService) Create(req request.CreateUserRequest) error {
	args := _m.Called(req)
	return args.Error(0)
}

func (_m *MockUserService) GetAll(page int, pageSize int) ([]response.UserResponse, error) {
	args := _m.Called(page, pageSize)
	return args.Get(0).([]response.UserResponse), args.Error(1)
}

func (_m *MockUserService) GetByID(userID uint) (*response.UserResponse, error) {
	args := _m.Called(userID)
	return args.Get(0).(*response.UserResponse), args.Error(1)
}

func (_m *MockUserService) Update(userID uint, req request.UpdateUserRequest) error {
	args := _m.Called(userID, req)
	return args.Error(0)
}

func (_m *MockUserService) Delete(userID uint) error {
	args := _m.Called(userID)
	return args.Error(0)
}
