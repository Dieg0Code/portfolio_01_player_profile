package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Achievement struct {
	gorm.Model
	Name            string        `gorm:"type:varchar(255);not null" validate:"required"`
	Description     string        `gorm:"type:varchar(255);not null" validate:"required"`
	PlayerProfileID uint          `gorm:"type:int;not null" validate:"required"` // Clave foránea
	PlayerProfile   PlayerProfile `gorm:"foreignKey:PlayerProfileID"`            // Relación con PlayerProfile
}

// Validate validates the Achievement struct.
func (a *Achievement) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
