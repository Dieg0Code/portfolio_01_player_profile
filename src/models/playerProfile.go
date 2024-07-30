package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PlayerProfile struct {
	gorm.Model
	Nickname     string        `gorm:"type:varchar(255);unique;not null" validate:"required"`
	Avatar       string        `gorm:"type:varchar(255);not null" validate:"required"`
	Level        int           `gorm:"type:int;not null" validate:"required"`
	Experience   int           `gorm:"type:int;not null" validate:"required"`
	Points       int           `gorm:"type:int;not null" validate:"required"`
	UserID       uint          `gorm:"type:int;not null" validate:"required"` // Clave foránea
	User         User          `gorm:"foreignKey:UserID"`                     // Relación con User
	Achievements []Achievement `gorm:"many2many:player_profile_achievements"`
}

func (p *PlayerProfile) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
