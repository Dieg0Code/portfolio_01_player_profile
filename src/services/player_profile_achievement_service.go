package services

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
)

type PlayerProfileAchievementService interface {
	Create(playerProfileAchievement request.CreatePlayerProfileAchievementRequest) error
	Delete(playerProfileID uint, achievementID uint) error
	GetAllPlayerProfileAchievements(offset int, pageSize int) ([]response.PlayerProfileAchievementResponse, error)
	GetByPlayerID(playerProfileID uint) ([]response.PlayerProfileAchievementResponse, error)
	GetByAchievementID(achievementID uint) ([]response.PlayerProfileAchievementResponse, error)
	GetByPlayerIDAndAchievementID(playerProfileID uint, achievementID uint) (*response.PlayerProfileAchievementResponse, error)
}
