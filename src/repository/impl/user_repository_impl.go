package impl

import (
	"errors"

	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	r "github.com/dieg0code/player-profile/src/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// FindByEmail implements repository.UserRepository.
func (u *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User

	result := u.Db.Where(EmailPlaceHolder, email).First(&user)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[UserRepositoryImpl.FindByEmail] Failed to find user")
		return nil, helpers.ErrorUserNotFound
	}
	return &user, nil
}

func NewUserRepositoryImpl(db *gorm.DB) r.UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// CreateUser implements repository.UserRepository.
func (u *UserRepositoryImpl) CreateUser(user *models.User) error {

	result := u.Db.Create(user)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[UserRepositoryImpl.CreateUser] Failed to create user")
		return result.Error
	}
	return nil
}

// GetAllUsers implements repository.UserRepository with pagination.
func (u *UserRepositoryImpl) GetAllUsers(offset int, pageSize int) ([]models.User, error) {
	var users []models.User

	result := u.Db.Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[UserRepositoryImpl.GetAllUsers] Failed to get all users")
		return nil, helpers.ErrorGetAllUsers
	}

	return users, nil
}

// GetUser implements repository.UserRepository.
func (u *UserRepositoryImpl) GetUser(userID uint) (*models.User, error) {

	exists, err := u.CheckUserExists(userID)
	if err != nil {
		logrus.WithError(err).Error("[UserRepositoryImpl.GetUser] Failed to check if user exists")
		return nil, err
	}

	if !exists {
		return nil, helpers.ErrorUserNotFound
	}

	var userFound models.User

	result := u.Db.Where(IDPlaceHolder, userID).First(&userFound)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("[UserRepositoryImpl.GetUser] Failed to get user")
		return nil, result.Error
	}

	// Return the found user.
	return &userFound, nil
}

// UpdateUser implements repository.UserRepository.
func (u *UserRepositoryImpl) UpdateUser(userID uint, user *models.User) error {

	exists, err := u.CheckUserExists(userID)
	if err != nil {
		logrus.WithError(err).Error("[UserRepositoryImpl.UpdateUser] Failed to check if user exists")
		return err
	}

	if !exists {
		return helpers.ErrorUserNotFound
	}

	result := u.Db.Model(&models.User{}).Where(IDPlaceHolder, userID).Updates(user)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[UserRepositoryImpl.UpdateUser] Failed to update user")
		return helpers.ErrorUpdateUser
	}

	return nil
}

func (u *UserRepositoryImpl) DeleteUser(userID uint) error {
	exists, err := u.CheckUserExists(userID)
	if err != nil {
		logrus.WithError(err).Error("[UserRepositoryImpl.DeleteUser] Failed to check if user exists")
		return err
	}
	if !exists {
		return helpers.ErrorUserNotFound
	}

	result := u.Db.Where(IDPlaceHolder, userID).Delete(&models.User{})
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[UserRepositoryImpl.DeleteUser] Failed to delete user")
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
		logrus.WithField("userID", userID).Error("[UserRepositoryImpl.CheckUserExists] User not found")
		return false, helpers.ErrorUserNotFound
	}
	// Other Kind of error
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[UserRepositoryImpl.CheckUserExists] Failed to check if user exists")
		return false, result.Error
	}

	return exists > 0, nil
}
