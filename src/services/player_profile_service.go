package services

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
)

type PlayerProfileService interface {
	Create(playerProfile request.CreatePlayerProfileRequest) error
	GetByID(playerProfileID uint) (*response.PlayerProfileResponse, error)
	Update(playerProfileID uint, playerProfile request.UpdatePlayerProfileRequest) error
	Delete(playerProfileID uint) error
	GetAll() ([]response.UserResponse, error)
}
