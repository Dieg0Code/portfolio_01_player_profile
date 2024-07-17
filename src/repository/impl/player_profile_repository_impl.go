package impl

import (
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/repository"
	"gorm.io/gorm"
)

type PlayerProfileRepositoryImpl struct {
	Db *gorm.DB
}

func NewPlayerProfileRepositoryImpl(db *gorm.DB) repository.PlayerProfileRepository {
	return &PlayerProfileRepositoryImpl{Db: db}
}

// CreatePlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) CreatePlayerProfile(playerProfile *models.PlayerProfile) error {
	panic("unimplemented")
}

// GetPlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) GetPlayerProfile(playerProfileID int) (*models.PlayerProfile, error) {
	panic("unimplemented")
}

// UpdatePlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) UpdatePlayerProfile(playerProfile *models.PlayerProfile) error {
	panic("unimplemented")
}

// DeletePlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) DeletePlayerProfile(playerProfileID int) error {
	panic("unimplemented")
}
