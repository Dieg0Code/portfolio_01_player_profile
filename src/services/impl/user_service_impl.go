package impl

import (
	"errors"
	"os"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/repository"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
	PasswordHasher services.PasswordHasher
}

// Create implements services.UserService.
func (u *UserServiceImpl) Create(user request.CreateUserRequest) error {

	err := u.Validate.Struct(user)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.Create] Failed to validate user data")
		return err
	}

	hashedPassword, err := u.PasswordHasher.HashPassword(user.Password)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.Create] Failed to hash password")
		return errors.New("failed to hash password")
	}

	userModel := models.User{
		UserName: user.UserName,
		PassWord: string(hashedPassword),
		Email:    user.Email,
		Age:      user.Age,
		Role:     os.Getenv("DEFAULT_ROLE"),
	}

	err = u.UserRepository.CreateUser(&userModel)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.Create] Failed to create user")
		return errors.New("email already exists")
	}

	return nil
}

// Delete implements services.UserService.
func (u *UserServiceImpl) Delete(userID uint) error {

	if userID == 0 {
		return helpers.ErrInvalidUserID
	}

	err := u.UserRepository.DeleteUser(userID)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.Delete] Failed to delete user")
		return err
	}

	return nil
}

// GetAll implements services.UserService.
func (u *UserServiceImpl) GetAll(page int, pageSize int) ([]response.UserResponse, error) {

	if page <= 0 || pageSize <= 0 {
		return nil, helpers.ErrInvalidPagination
	}

	offset := (page - 1) * pageSize

	users, err := u.UserRepository.GetAllUsers(offset, pageSize)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.GetAll] Failed to get all users")
		return nil, err
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponse := response.UserResponse{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
			Age:      user.Age,
		}

		err = u.Validate.Struct(userResponse)
		if err != nil {
			logrus.WithError(err).Error("[UserServiceImpl.GetAll] Failed to validate user data")
			return nil, helpers.ErrUserDataValidation
		}

		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

// GetByID implements services.UserService.
func (u *UserServiceImpl) GetByID(userID uint) (*response.UserResponse, error) {

	if userID == 0 {
		return nil, helpers.ErrInvalidUserID
	}

	user, err := u.UserRepository.GetUser(userID)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.GetByID] Failed to get user")
		return nil, err
	}

	if user == nil {
		return nil, helpers.ErrorUserNotFound
	}

	userResponse := response.UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Age:      user.Age,
	}

	err = u.Validate.Struct(userResponse)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.GetByID] Failed to validate user data")
		return nil, helpers.ErrUserDataValidation
	}

	return &userResponse, nil
}

// Update implements services.UserService.
func (u *UserServiceImpl) Update(userID uint, user request.UpdateUserRequest) error {

	userData, err := u.UserRepository.GetUser(userID)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.Update] Failed to get user")
		return err
	}

	err = u.Validate.Struct(user)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.Update] Failed to validate user data")
		return err
	}

	userData.UserName = user.UserName
	userData.Email = user.Email
	userData.Age = user.Age

	err = u.UserRepository.UpdateUser(userID, userData)
	if err != nil {
		logrus.WithError(err).Error("[UserServiceImpl.Update] Failed to update user")
		return err
	}

	return nil
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate, passwordHasher services.PasswordHasher) services.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
		PasswordHasher: passwordHasher,
	}
}
