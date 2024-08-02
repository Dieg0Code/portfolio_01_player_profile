package mocks

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/stretchr/testify/mock"
)

type MockPlayerProfileService struct {
	mock.Mock
}

func (_m *MockPlayerProfileService) Create(playerProfile request.CreatePlayerProfileRequest) error {
	args := _m.Called(playerProfile)
	return args.Error(0)
}
func (_m *MockPlayerProfileService) GetByID(playerProfileID uint) (*response.PlayerProfileResponse, error) {
	args := _m.Called(playerProfileID)
	return args.Get(0).(*response.PlayerProfileResponse), args.Error(1)
}
func (_m *MockPlayerProfileService) Update(playerProfileID uint, playerProfile request.UpdatePlayerProfileRequest) error {
	args := _m.Called(playerProfileID, playerProfile)
	return args.Error(0)
}
func (_m *MockPlayerProfileService) Delete(playerProfileID uint) error {
	args := _m.Called(playerProfileID)
	return args.Error(0)
}
func (_m *MockPlayerProfileService) GetAll(page int, pageSize int) ([]response.PlayerProfileResponse, error) {
	args := _m.Called(page, pageSize)
	return args.Get(0).([]response.PlayerProfileResponse), args.Error(1)
}

func (_m *MockPlayerProfileService) GetPlayerWithAchievements(playerProfileID uint) (*response.PlayerWithAchievements, error) {
	args := _m.Called(playerProfileID)
	return args.Get(0).(*response.PlayerWithAchievements), args.Error(1)
}
