package impl

import (
	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	r "github.com/dieg0code/player-profile/src/repository"
	"gorm.io/gorm"
)

type PlayerProfileAchievementRepositoryImpl struct {
	db *gorm.DB
}

// FindByPlayerProfileIDAndAchievementID implements repository.PlayerProfileAchievementRepository.
func (p *PlayerProfileAchievementRepositoryImpl) FindByPlayerProfileIDAndAchievementID(playerProfileID uint, achievementID uint) (*models.PlayerProfileAchievement, error) {
	var playerProfileAchievement models.PlayerProfileAchievement
	result := p.db.Model(&models.PlayerProfileAchievement{}).Where(PlayerAndAchievementIDPlaceHolder, playerProfileID, achievementID).First(&playerProfileAchievement)
	if result.Error != nil {
		return nil, result.Error
	}

	return &playerProfileAchievement, nil
}

// Delete implements repository.PlayerProfileAchievementRepository.
func (p *PlayerProfileAchievementRepositoryImpl) Delete(profileID uint, achievementID uint) error {
	result := p.db.Where(PlayerAndAchievementIDPlaceHolder, profileID, achievementID).Delete(&models.PlayerProfileAchievement{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllPlayerProfileAchievements implements repository.PlayerProfileAchievementRepository.
func (p *PlayerProfileAchievementRepositoryImpl) GetAllPlayerProfileAchievements(offset int, pageSize int) ([]models.PlayerProfileAchievement, error) {
	var playerProfileAchievements []models.PlayerProfileAchievement
	result := p.db.Model(&models.PlayerProfileAchievement{}).Offset(offset).Limit(pageSize).Find(&playerProfileAchievements)
	if result.Error != nil {
		return nil, result.Error
	}

	return playerProfileAchievements, nil
}

// CheckPlayerProfileAchievementExists implements repository.PlayerProfileAchievementRepository.
func (p *PlayerProfileAchievementRepositoryImpl) CheckPlayerProfileAchievementExists(playerProfileAchieventID uint) (bool, error) {
	var exists int64

	result := p.db.Model(&models.PlayerProfileAchievement{}).Where(IDPlaceHolder, playerProfileAchieventID).Count(&exists)
	if result.Error != nil {
		return false, result.Error
	}

	return exists > 0, nil
}

// Create implements repository.PlayerProfileAchievementRepository.
func (p *PlayerProfileAchievementRepositoryImpl) Create(playerProfileAchievement *models.PlayerProfileAchievement) error {
	result := p.db.Create(playerProfileAchievement)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindByAchievementID implements repository.PlayerProfileAchievementRepository.
func (p *PlayerProfileAchievementRepositoryImpl) FindByAchievementID(achievementID uint) ([]models.PlayerProfileAchievement, error) {
	exists, err := p.CheckPlayerProfileAchievementExists(achievementID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, helpers.ErrRegisterNotFound
	}

	var playerProfileAchievements []models.PlayerProfileAchievement
	result := p.db.Model(&models.PlayerProfileAchievement{}).Where(AchievementIDPlaceHolder, achievementID).Find(&playerProfileAchievements)
	if result.Error != nil {
		return nil, result.Error
	}

	return playerProfileAchievements, nil
}

// FindByPlayerProfileID implements repository.PlayerProfileAchievementRepository.
func (p *PlayerProfileAchievementRepositoryImpl) FindByPlayerProfileID(playerProfileID uint) ([]models.PlayerProfileAchievement, error) {
	exists, err := p.CheckPlayerProfileAchievementExists(playerProfileID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, helpers.ErrRegisterNotFound
	}

	var playerProfileAchievements []models.PlayerProfileAchievement
	result := p.db.Model(&models.PlayerProfileAchievement{}).Where(PlayerProfileIDPlaceHolder, playerProfileID).Find(&playerProfileAchievements)
	if result.Error != nil {
		return nil, result.Error
	}

	return playerProfileAchievements, nil
}

func NewPlayerProfileAchievementRepositoryImpl(db *gorm.DB) r.PlayerProfileAchievementRepository {
	return &PlayerProfileAchievementRepositoryImpl{db: db}
}
