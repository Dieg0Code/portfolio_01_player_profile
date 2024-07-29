package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string          `gorm:"type:varchar(255);unique;not null" validate:"required"`
	PassWord string          `gorm:"type:varchar(255);not null" validate:"required"`
	Email    string          `gorm:"type:varchar(255);unique;not null" validate:"required"`
	Age      int             `gorm:"type:int;not null" validate:"required"`
	Role     string          `gorm:"type:varchar(255);not null" validate:"required,oneof=admin user"`
	Profiles []PlayerProfile `gorm:"foreignKey:UserID"` // Relación uno a muchos con PlayerProfile
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
