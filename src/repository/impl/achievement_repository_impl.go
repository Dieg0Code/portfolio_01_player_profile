package impl

import (
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/repository"
	"gorm.io/gorm"
)

type AchivementRepositoryImpl struct {
	Db *gorm.DB
}

// CreateAchievement implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) CreateAchievement(achievement *models.Achievement) error {
	panic("unimplemented")
}

// DeleteAchievement implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) DeleteAchievement(achievementID int) error {
	panic("unimplemented")
}

// GetAchievement implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) GetAchievement(achievementID int) (*models.Achievement, error) {
	panic("unimplemented")
}

// UpdateAchievement implements repository.AchievementRepository.
func (a *AchivementRepositoryImpl) UpdateAchievement(achievement *models.Achievement) error {
	panic("unimplemented")
}

func NewAchievementRepositoryImpl(db *gorm.DB) repository.AchievementRepository {
	return &AchivementRepositoryImpl{Db: db}
}
