package repository

import "github.com/dieg0code/player-profile/src/models"

type PlayerProfileAchievementRepository interface {
	Create(playerProfileAchievement *models.PlayerProfileAchievement) error
	FindByPlayerProfileID(playerProfileID uint) ([]models.PlayerProfileAchievement, error)
	FindByAchievementID(achievementID uint) ([]models.PlayerProfileAchievement, error)
	FindByPlayerProfileIDAndAchievementID(playerProfileID, achievementID uint) (*models.PlayerProfileAchievement, error)
	Delete(profileID, achievementID uint) error
	CheckPlayerProfileAchievementExists(playerProfileAchieventID uint) (bool, error)
	GetAllPlayerProfileAchievements(offset int, pageSize int) ([]models.PlayerProfileAchievement, error)
}
