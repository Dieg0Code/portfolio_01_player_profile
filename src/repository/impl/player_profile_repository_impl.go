package impl

import (
	"errors"

	"github.com/dieg0code/player-profile/src/helpers"
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
	result := p.Db.Create(playerProfile)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetPlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) GetPlayerProfile(playerProfileID int) (*models.PlayerProfile, error) {
	exists, err := p.CheckPlayerProfileExists(playerProfileID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, helpers.ErrorPlayerProfileNotFound
	}

	var playerProfileFound models.PlayerProfile

	result := p.Db.Where("player_profile_id = ?", playerProfileID).First(&playerProfileFound)

	if result.Error != nil {
		return nil, result.Error
	}

	return &playerProfileFound, nil
}

// UpdatePlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) UpdatePlayerProfile(playerProfile *models.PlayerProfile) error {
	exists, err := p.CheckPlayerProfileExists(playerProfile.PlayerProfileID)
	if err != nil {
		return err
	}

	if !exists {
		return helpers.ErrorPlayerProfileNotFound
	}

	result := p.Db.Model(&models.PlayerProfile{}).Where("player_profile_id = ?", playerProfile.PlayerProfileID).Updates(playerProfile)
	if result.Error != nil {
		return helpers.ErrorUpdatePlayer
	}

	return nil
}

// DeletePlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) DeletePlayerProfile(playerProfileID int) error {
	exists, err := p.CheckPlayerProfileExists(playerProfileID)

	if err != nil {
		return err
	}

	if !exists {
		return helpers.ErrorPlayerProfileNotFound
	}

	result := p.Db.Delete(&models.PlayerProfile{}, playerProfileID)

	if result.Error != nil {
		return helpers.ErrorDeletingUser
	}

	return nil
}

// Check if Player Profile exists.
func (p *PlayerProfileRepositoryImpl) CheckPlayerProfileExists(playerProfileID int) (bool, error) {
	var exists int64

	result := p.Db.Model(&models.PlayerProfile{}).Where("player_profile_id = ?", playerProfileID).Count(&exists)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, helpers.ErrorPlayerProfileNotFound
	}

	// Other Kind of error
	if result.Error != nil {
		return false, result.Error
	}
	return exists > 0, nil
}