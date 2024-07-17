package impl

import (
	"errors"

	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	repo "github.com/dieg0code/player-profile/src/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repo.UserRepository {
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
func (u *UserRepositoryImpl) GetUser(userID int) (*models.User, error) {

	exists, err := u.CheckUserExists(userID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, helpers.ErrorUserNotFound
	}

	var userFound models.User

	// Explicitly specify the field name in the query.
	result := u.Db.Where("user_id = ?", userID).First(&userFound)

	// Check for a "record not found" error.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, helpers.ErrorUserNotFound
	}

	// Check for other types of errors.
	if result.Error != nil {
		return nil, result.Error
	}

	// Return the found user.
	return &userFound, nil
}

// UpdateUser implements repository.UserRepository.
func (u *UserRepositoryImpl) UpdateUser(user *models.User) error {
	// Optimización: Verificar existencia sin cargar el usuario completo.
	exists, err := u.CheckUserExists(user.UserID)
	if err != nil {
		return err
	}

	if !exists {
		return helpers.ErrorUserNotFound
	}

	result := u.Db.Model(&models.User{}).Where("user_id = ?", user.UserID).Updates(user)
	if result.Error != nil {
		return helpers.ErrorUpdateUser
	}
	return nil
}

func (u *UserRepositoryImpl) DeleteUser(userID int) error {
	exists, err := u.CheckUserExists(userID)
	if err != nil {
		return err
	}
	if !exists {
		return helpers.ErrorUserNotFound
	}

	result := u.Db.Delete(&models.User{}, userID)
	if result.Error != nil {
		return helpers.ErrorDeleteUser
	}
	return nil
}

// CheckUserExists verifica si existe un usuario con el ID proporcionado.
func (u *UserRepositoryImpl) CheckUserExists(userID int) (bool, error) {
	var exists int64
	result := u.Db.Model(&models.User{}).Where("user_id = ?", userID).Count(&exists)
	if result.Error != nil {
		return false, result.Error
	}
	return exists > 0, nil
}
