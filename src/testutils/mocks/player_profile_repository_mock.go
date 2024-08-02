package mocks

import (
	"github.com/dieg0code/player-profile/src/models"
	"github.com/stretchr/testify/mock"
)

type PlayerProfileRepository struct {
	mock.Mock
}

func (_m *PlayerProfileRepository) CreatePlayerProfile(playerProfile *models.PlayerProfile) error {
	ret := _m.Called(playerProfile)
	return ret.Error(0)
}

func (_m *PlayerProfileRepository) GetPlayerProfile(playerProfileID uint) (*models.PlayerProfile, error) {
	args := _m.Called(playerProfileID)

	player, _ := args.Get(0).(*models.PlayerProfile)

	return player, args.Error(1)
}

func (_m *PlayerProfileRepository) GetAllPlayerProfiles(offset int, pageSize int) ([]models.PlayerProfile, error) {
	ret := _m.Called(offset, pageSize)
	return ret.Get(0).([]models.PlayerProfile), ret.Error(1)
}

func (_m *PlayerProfileRepository) UpdatePlayerProfile(playerProfileID uint, playerProfile *models.PlayerProfile) error {
	ret := _m.Called(playerProfileID, playerProfile)
	return ret.Error(0)
}

func (_m *PlayerProfileRepository) DeletePlayerProfile(playerProfileID uint) error {
	ret := _m.Called(playerProfileID)
	return ret.Error(0)
}

func (_m *PlayerProfileRepository) CheckPlayerProfileExists(playerProfileID uint) (bool, error) {
	args := _m.Called(playerProfileID)
	return args.Bool(0), args.Error(1)
}

func (_m *PlayerProfileRepository) GetPlayerWithAchievements(playerProfileID uint) (*models.PlayerProfile, error) {
	args := _m.Called(playerProfileID)

	player, _ := args.Get(0).(*models.PlayerProfile)

	return player, args.Error(1)
}
