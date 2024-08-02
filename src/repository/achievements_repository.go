package repository

import "github.com/dieg0code/player-profile/src/models"

type AchievementRepository interface {
	CreateAchievement(achievement *models.Achievement) error
	GetAchievement(achievementID uint) (*models.Achievement, error)
	UpdateAchievement(achievementID uint, achievement *models.Achievement) error
	DeleteAchievement(achievementID uint) error
	CheckAchievementExists(achievementID uint) (bool, error)
	GetAllAchievements(offset int, pageSize int) ([]models.Achievement, error)
	GetAchievementWithPlayers(achievementID uint) (*models.Achievement, error)
}
