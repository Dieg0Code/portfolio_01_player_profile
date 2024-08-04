package impl

import (
	"errors"

	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	r "github.com/dieg0code/player-profile/src/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PlayerProfileRepositoryImpl struct {
	Db *gorm.DB
}

// GetPlayerWithAchievements implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) GetPlayerWithAchievements(playerProfileID uint) (*models.PlayerProfile, error) {
	exists, err := p.CheckPlayerProfileExists(playerProfileID)
	if err != nil {
		logrus.WithError(err).Error("[PlayerProfileRepositoryImpl.GetPlayerWithAchievements] Failed to check if player profile exists")
		return nil, err
	}

	if !exists {
		return nil, helpers.ErrorPlayerProfileNotFound
	}

	var playerProfileFound models.PlayerProfile

	result := p.Db.Preload("Achievements").Where(IDPlaceHolder, playerProfileID).First(&playerProfileFound)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("[PlayerProfileRepositoryImpl.GetPlayerWithAchievements] Failed to get player profile with achievements")
		return nil, result.Error
	}

	return &playerProfileFound, nil
}

func NewPlayerProfileRepositoryImpl(db *gorm.DB) r.PlayerProfileRepository {
	return &PlayerProfileRepositoryImpl{Db: db}
}

// CreatePlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) CreatePlayerProfile(playerProfile *models.PlayerProfile) error {

	result := p.Db.Create(playerProfile)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[PlayerProfileRepositoryImpl.CreatePlayerProfile] Failed to create player profile")
		return result.Error
	}

	return nil
}

// GetPlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) GetPlayerProfile(playerProfileID uint) (*models.PlayerProfile, error) {
	exists, err := p.CheckPlayerProfileExists(playerProfileID)
	if err != nil {
		logrus.WithError(err).Error("[PlayerProfileRepositoryImpl.GetPlayerProfile] Failed to check if player profile exists")
		return nil, err
	}

	if !exists {
		return nil, helpers.ErrorPlayerProfileNotFound
	}

	var playerProfileFound models.PlayerProfile

	result := p.Db.Where(IDPlaceHolder, playerProfileID).First(&playerProfileFound)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("[PlayerProfileRepositoryImpl.GetPlayerProfile] Failed to get player profile")
		return nil, result.Error
	}

	return &playerProfileFound, nil
}

// GetAllPlayerProfiles implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) GetAllPlayerProfiles(offset int, pageSize int) ([]models.PlayerProfile, error) {
	var playerProfiles []models.PlayerProfile

	result := p.Db.Offset(offset).Limit(pageSize).Find(&playerProfiles)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[PlayerProfileRepositoryImpl.GetAllPlayerProfiles] Failed to get all player profiles")
		return nil, helpers.ErrorGetAllPlayerProfiles
	}

	return playerProfiles, nil
}

// UpdatePlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) UpdatePlayerProfile(playerProfileID uint, playerProfile *models.PlayerProfile) error {
	exists, err := p.CheckPlayerProfileExists(playerProfileID)
	if err != nil {
		logrus.WithError(err).Error("[PlayerProfileRepositoryImpl.UpdatePlayerProfile] Failed to check if player profile exists")
		return err
	}

	if !exists {
		return helpers.ErrorPlayerProfileNotFound
	}

	result := p.Db.Model(&models.PlayerProfile{}).Where(IDPlaceHolder, playerProfileID).Updates(playerProfile)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[PlayerProfileRepositoryImpl.UpdatePlayerProfile] Failed to update player profile")
		return helpers.ErrorUpdatePlayer
	}

	return nil
}

// DeletePlayerProfile implements repository.PlayerProfileRepository.
func (p *PlayerProfileRepositoryImpl) DeletePlayerProfile(playerProfileID uint) error {
	exists, err := p.CheckPlayerProfileExists(playerProfileID)

	if err != nil {
		logrus.WithError(err).Error("[PlayerProfileRepositoryImpl.DeletePlayerProfile] Failed to check if player profile exists")
		return err
	}

	if !exists {
		return helpers.ErrorPlayerProfileNotFound
	}

	result := p.Db.Delete(&models.PlayerProfile{}, playerProfileID)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("[PlayerProfileRepositoryImpl.DeletePlayerProfile] Failed to delete player profile")
		return helpers.ErrorDeletingUser
	}

	return nil
}

// Check if Player Profile exists.
func (p *PlayerProfileRepositoryImpl) CheckPlayerProfileExists(playerProfileID uint) (bool, error) {
	var exists int64

	result := p.Db.Model(&models.PlayerProfile{}).Where(IDPlaceHolder, playerProfileID).Count(&exists)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logrus.WithField("playerProfileID", playerProfileID).Error("[PlayerProfileRepositoryImpl.CheckPlayerProfileExists] Player Profile not found")
		return false, helpers.ErrorPlayerProfileNotFound
	}

	// Other Kind of error
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[PlayerProfileRepositoryImpl.CheckPlayerProfileExists] Failed to check if player profile exists")
		return false, result.Error
	}
	return exists > 0, nil
}
