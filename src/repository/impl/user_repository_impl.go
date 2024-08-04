package impl

import (
	"errors"

	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	r "github.com/dieg0code/player-profile/src/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func init() {
	// Configurar Logrus
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// FindByEmail implements repository.UserRepository.
func (u *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	logrus.Info("[UserRepositoryImpl.FindByEmail] Finding User: ", email)
	var user models.User

	result := u.Db.Where(EmailPlaceHolder, email).First(&user)
	if result.Error != nil {
		logrus.Error("[UserRepositoryImpl.FindByEmail] Error: ", result.Error)
		return nil, helpers.ErrorUserNotFound
	}

	logrus.Info("[UserRepositoryImpl.FindByEmail] User finded")
	return &user, nil
}

func NewUserRepositoryImpl(db *gorm.DB) r.UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// CreateUser implements repository.UserRepository.
func (u *UserRepositoryImpl) CreateUser(user *models.User) error {
	logger := logrus.WithFields(logrus.Fields{
		"function": "UserRepositoryImpl.CreateUser",
		"user":     user.Email,
	})

	result := u.Db.Create(user)
	if result.Error != nil {
		logger.WithError(result.Error).Error("Failed to create user")
		return result.Error
	}

	logger.Info("User created successfully")
	return nil
}

// GetAllUsers implements repository.UserRepository with pagination.
func (u *UserRepositoryImpl) GetAllUsers(offset int, pageSize int) ([]models.User, error) {
	logger := logrus.WithFields(logrus.Fields{
		"function": "UserRepositoryImpl.GetAllUsers",
		"offset":   offset,
		"pageSize": pageSize,
	})
	var users []models.User

	result := u.Db.Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		logger.WithError(result.Error).Error("Failed to get all users")
		return nil, helpers.ErrorGetAllUsers
	}

	logger.Info("Users retrieved successfully")
	return users, nil
}

// GetUser implements repository.UserRepository.
func (u *UserRepositoryImpl) GetUser(userID uint) (*models.User, error) {
	logger := logrus.WithFields(logrus.Fields{
		"function": "UserRepositoryImpl.GetUser",
		"userID":   userID,
	})

	exists, err := u.CheckUserExists(userID)
	if err != nil {
		logger.WithError(err).Error("Failed to check if user exists")
		return nil, err
	}

	if !exists {
		logrus.Error("[UserRepositoryImpl.GetUser] Error: ", helpers.ErrorUserNotFound)
		return nil, helpers.ErrorUserNotFound
	}

	var userFound models.User

	// Explicitly specify the field name in the query.
	result := u.Db.Where(IDPlaceHolder, userID).First(&userFound)

	// Check for other types of errors.
	if result.Error != nil {
		logrus.Error("[UserRepositoryImpl.GetUser] Error: ", result.Error)
		return nil, result.Error
	}

	// Return the found user.
	return &userFound, nil
}

// UpdateUser implements repository.UserRepository.
func (u *UserRepositoryImpl) UpdateUser(userID uint, user *models.User) error {
	logger := logrus.WithFields(logrus.Fields{
		"function":   "UserRepositoryImpl.UpdateUser",
		"userID":     userID,
		"user_email": user.Email,
	})
	exists, err := u.CheckUserExists(userID)
	if err != nil {
		logger.WithError(err).Error("Failed to check if user exists")
		return err
	}

	if !exists {
		return helpers.ErrorUserNotFound
	}

	result := u.Db.Model(&models.User{}).Where(IDPlaceHolder, userID).Updates(user)
	if result.Error != nil {
		logger.WithError(result.Error).Error("Failed to update user")
		return helpers.ErrorUpdateUser
	}

	logger.Info("User updated successfully")
	return nil
}

func (u *UserRepositoryImpl) DeleteUser(userID uint) error {
	exists, err := u.CheckUserExists(userID)
	if err != nil {
		logrus.Error("[UserRepositoryImpl.DeleteUser] Error: ", err)
		return err
	}
	if !exists {
		return helpers.ErrorUserNotFound
	}

	result := u.Db.Where(IDPlaceHolder, userID).Delete(&models.User{})
	if result.Error != nil {
		logrus.Error("[UserRepositoryImpl.DeleteUser] Error: ", result.Error)
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
		logrus.Error("[UserRepositoryImpl.CheckUserExists] Error: ", helpers.ErrorUserNotFound)
		return false, helpers.ErrorUserNotFound
	}
	// Other Kind of error
	if result.Error != nil {
		logrus.Error("[UserRepositoryImpl.CheckUserExists] Error: ", result.Error)
		return false, result.Error
	}

	return exists > 0, nil
}
