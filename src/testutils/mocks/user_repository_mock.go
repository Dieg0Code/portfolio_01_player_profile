package mocks

import (
	"github.com/dieg0code/player-profile/src/models"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (_m *UserRepository) CreateUser(user *models.User) error {
	ret := _m.Called(user)
	return ret.Error(0)
}

func (_m *UserRepository) GetUser(userID uint) (*models.User, error) {
	args := _m.Called(userID)

	user, _ := args.Get(0).(*models.User)

	return user, args.Error(1)
}

func (_m *UserRepository) GetAllUsers(offset int, pageSize int) ([]models.User, error) {
	ret := _m.Called(offset, pageSize)
	return ret.Get(0).([]models.User), ret.Error(1)
}

func (_m *UserRepository) UpdateUser(userID uint, user *models.User) error {
	ret := _m.Called(userID, user)
	return ret.Error(0)
}

func (_m *UserRepository) DeleteUser(userID uint) error {
	ret := _m.Called(userID)
	return ret.Error(0)
}

func (_m *UserRepository) FindByEmail(email string) (*models.User, error) {
	args := _m.Called(email)

	user, _ := args.Get(0).(*models.User)

	return user, args.Error(1)
}
