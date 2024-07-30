package models

import "gorm.io/gorm"

type PlayerProfileAchievement struct {
	gorm.Model
	PlayerProfileID uint `gorm:"type:int;not null" validate:"required"`
	AchievementID   uint `gorm:"type:int;not null" validate:"required"`
}
