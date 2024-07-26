package mocks

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/stretchr/testify/mock"
)

type MockAchievementService struct {
	mock.Mock
}

func (_m *MockAchievementService) Create(achievement request.CreateAchievementRequest) error {
	ret := _m.Called(achievement)
	return ret.Error(0)
}
func (_m *MockAchievementService) Delete(achievementID uint) error {
	ret := _m.Called(achievementID)
	return ret.Error(0)
}
func (_m *MockAchievementService) GetByID(achievementID uint) (*response.AchievementResponse, error) {
	ret := _m.Called(achievementID)
	return ret.Get(0).(*response.AchievementResponse), ret.Error(1)
}
func (_m *MockAchievementService) GetAll(page int, pageSize int) ([]response.AchievementResponse, error) {
	ret := _m.Called(page, pageSize)
	return ret.Get(0).([]response.AchievementResponse), ret.Error(1)
}
func (_m *MockAchievementService) Update(achievementID uint, achievement request.UpdateAchievementRequest) error {
	ret := _m.Called(achievementID, achievement)
	return ret.Error(0)
}
