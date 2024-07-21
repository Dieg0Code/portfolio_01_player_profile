package repository

import "github.com/dieg0code/player-profile/src/models"

type PlayerProfileRepository interface {
	CreatePlayerProfile(playerProfile *models.PlayerProfile) error
	GetPlayerProfile(playerProfileID uint) (*models.PlayerProfile, error)
	UpdatePlayerProfile(playerProfile *models.PlayerProfile) error
	DeletePlayerProfile(playerProfileID uint) error
	CheckPlayerProfileExists(playerProfileID uint) (bool, error)
}
