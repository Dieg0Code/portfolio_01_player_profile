package repository

import "github.com/dieg0code/player-profile/src/models"

type AchievementRepository interface {
	CreateAchievement(achievement *models.Achievement) error
	GetAchievement(achievementID int) (*models.Achievement, error)
	UpdateAchievement(achievement *models.Achievement) error
	DeleteAchievement(achievementID int) error
}
