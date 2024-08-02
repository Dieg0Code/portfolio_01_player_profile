package impl

import (
	"errors"

	h "github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	r "github.com/dieg0code/player-profile/src/repository"
	"gorm.io/gorm"
)

type AchivementRepositoryImpl struct {
	Db *gorm.DB
}

// GetAchievementWithPlayers implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) GetAchievementWithPlayers(achievementID uint) (*models.Achievement, error) {
	var achievement models.Achievement

	result := a.Db.Preload("PlayerProfiles").Where(IDPlaceHolder, achievementID).First(&achievement)

	if result.Error != nil {
		return nil, result.Error
	}

	return &achievement, nil
}

func NewAchievementRepositoryImpl(db *gorm.DB) r.AchievementRepository {
	return &AchivementRepositoryImpl{Db: db}
}

// CreateAchievement implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) CreateAchievement(achievement *models.Achievement) error {

	result := a.Db.Create(achievement)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAchievement implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) GetAchievement(achievementID uint) (*models.Achievement, error) {
	exists, err := a.CheckAchievementExists(achievementID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, h.ErrorAchievementNotFound
	}

	var achievementFound models.Achievement

	result := a.Db.Where(IDPlaceHolder, achievementID).First(&achievementFound)

	if result.Error != nil {
		return nil, result.Error
	}

	return &achievementFound, nil
}

// GetAllAchievements implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) GetAllAchievements(offset int, pageSize int) ([]models.Achievement, error) {
	var achievements []models.Achievement

	result := a.Db.Offset(offset).Limit(pageSize).Find(&achievements)

	if result.Error != nil {
		return nil, result.Error
	}

	return achievements, nil
}

// UpdateAchievement implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) UpdateAchievement(achievementID uint, achievement *models.Achievement) error {
	exists, err := a.CheckAchievementExists(achievementID)
	if err != nil {
		return err
	}

	if !exists {
		return h.ErrorAchievementNotFound
	}

	result := a.Db.Model(&models.Achievement{}).Where(IDPlaceHolder, achievementID).Updates(achievement)
	if result.Error != nil {
		return h.ErrorUpdateAchievement
	}

	return nil
}

// DeleteAchievement implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) DeleteAchievement(achievementID uint) error {
	exists, err := a.CheckAchievementExists(achievementID)
	if err != nil {
		return err
	}

	if !exists {
		return h.ErrorAchievementNotFound
	}

	result := a.Db.Delete(&models.Achievement{}, IDPlaceHolder, achievementID)
	if result.Error != nil {
		return h.ErrorDeletingAchievement
	}

	return nil
}

// CheackAchievementExists implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) CheckAchievementExists(achievementID uint) (bool, error) {
	var exists int64

	result := a.Db.Model(&models.Achievement{}).Where(IDPlaceHolder, achievementID).Count(&exists)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, h.ErrorAchievementNotFound
	}

	// Other Kind of error
	if result.Error != nil {
		return false, result.Error
	}
	return exists > 0, nil
}
