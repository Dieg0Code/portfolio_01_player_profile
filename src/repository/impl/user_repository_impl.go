package impl

import (
	"errors"

	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	r "github.com/dieg0code/player-profile/src/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) r.UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// CreateUser implements repository.UserRepository.
func (u *UserRepositoryImpl) CreateUser(user *models.User) error {

	result := u.Db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetUser implements repository.UserRepository.
func (u *UserRepositoryImpl) GetUser(userID uint) (*models.User, error) {

	exists, err := u.CheckUserExists(userID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, helpers.ErrorUserNotFound
	}

	var userFound models.User

	// Explicitly specify the field name in the query.
	result := u.Db.Where(IDPlaceHolder, userID).First(&userFound)

	// Check for other types of errors.
	if result.Error != nil {
		return nil, result.Error
	}

	// Return the found user.
	return &userFound, nil
}

// UpdateUser implements repository.UserRepository.
func (u *UserRepositoryImpl) UpdateUser(user *models.User) error {
	exists, err := u.CheckUserExists(user.ID)
	if err != nil {
		return err
	}

	if !exists {
		return helpers.ErrorUserNotFound
	}

	result := u.Db.Model(&models.User{}).Where(IDPlaceHolder, user.ID).Updates(user)
	if result.Error != nil {
		return helpers.ErrorUpdateUser
	}
	return nil
}

func (u *UserRepositoryImpl) DeleteUser(userID uint) error {
	exists, err := u.CheckUserExists(userID)
	if err != nil {
		return err
	}
	if !exists {
		return helpers.ErrorUserNotFound
	}

	result := u.Db.Where(IDPlaceHolder, userID).Delete(&models.User{})
	if result.Error != nil {
		return helpers.ErrorDeleteUser
	}
	return nil
}

// CheckUserExists verifica si existe un usuario con el ID proporcionado.
func (u *UserRepositoryImpl) CheckUserExists(userID uint) (bool, error) {
	var exists int64

	result := u.Db.Model(&models.User{}).Where(IDPlaceHolder, userID).Count(&exists)

	// Check for a "record not found" error.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, helpers.ErrorUserNotFound
	}
	// Other Kind of error
	if result.Error != nil {
		return false, result.Error
	}

	return exists > 0, nil
}
