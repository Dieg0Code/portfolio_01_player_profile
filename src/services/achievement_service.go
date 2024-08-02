package services

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
)

type AchievementService interface {
	Create(achievement request.CreateAchievementRequest) error
	Delete(achievementID uint) error
	GetByID(achievementID uint) (*response.AchievementResponse, error)
	GetAll(page int, pageSize int) ([]response.AchievementResponse, error)
	Update(achievementID uint, achievement request.UpdateAchievementRequest) error
	GetAchievementWithPlayers(achievementID uint) (*response.AchievementWithPlayers, error)
}
