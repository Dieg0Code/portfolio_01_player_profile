package repository

import "github.com/dieg0code/player-profile/src/models"

type PlayerProfileRepository interface {
	CreatePlayerProfile(playerProfile *models.PlayerProfile) error
	GetPlayerProfile(playerProfileID int) (*models.PlayerProfile, error)
	UpdatePlayerProfile(playerProfile *models.PlayerProfile) error
	DeletePlayerProfile(playerProfileID int) error
	CheckPlayerProfileExists(playerProfileID int) (bool, error)
}
