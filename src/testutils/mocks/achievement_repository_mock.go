package mocks

import (
	"github.com/dieg0code/player-profile/src/models"
	"github.com/stretchr/testify/mock"
)

type AchievementRepository struct {
	mock.Mock
}

func (_m *AchievementRepository) CreateAchievement(achievement *models.Achievement) error {
	ret := _m.Called(achievement)
	return ret.Error(0)
}

func (_m *AchievementRepository) GetAchievement(achievementID uint) (*models.Achievement, error) {
	args := _m.Called(achievementID)

	achievement, _ := args.Get(0).(*models.Achievement)

	return achievement, args.Error(1)
}

func (_m *AchievementRepository) GetAllAchievements(offset int, pageSize int) ([]models.Achievement, error) {
	ret := _m.Called(offset, pageSize)
	return ret.Get(0).([]models.Achievement), ret.Error(1)
}

func (_m *AchievementRepository) UpdateAchievement(achievementID uint, achievement *models.Achievement) error {
	ret := _m.Called(achievementID, achievement)
	return ret.Error(0)
}

func (_m *AchievementRepository) DeleteAchievement(achievementID uint) error {
	ret := _m.Called(achievementID)
	return ret.Error(0)
}

func (_m *AchievementRepository) CheckAchievementExists(achievementID uint) (bool, error) {
	args := _m.Called(achievementID)
	return args.Bool(0), args.Error(1)
}

func (_m *AchievementRepository) GetAchievementWithPlayers(achievementID uint) (*models.Achievement, error) {
	args := _m.Called(achievementID)

	achievement, _ := args.Get(0).(*models.Achievement)

	return achievement, args.Error(1)
}
