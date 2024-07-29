package impl

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/repository"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/go-playground/validator/v10"
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
		return err
	}

	hashedPassword, err := u.PasswordHasher.HashPassword(user.Password)
	if err != nil {
		return err
	}

	userModel := models.User{
		UserName: user.UserName,
		PassWord: string(hashedPassword),
		Email:    user.Email,
		Age:      user.Age,
		Role:     "user",
	}

	err = u.UserRepository.CreateUser(&userModel)
	if err != nil {
		return err
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
		return nil, helpers.ErrUserDataValidation
	}

	return &userResponse, nil
}

// Update implements services.UserService.
func (u *UserServiceImpl) Update(userID uint, user request.UpdateUserRequest) error {

	userData, err := u.UserRepository.GetUser(userID)
	if err != nil {
		return err
	}

	err = u.Validate.Struct(user)
	if err != nil {
		return err
	}

	userData.UserName = user.UserName
	userData.Email = user.Email
	userData.Age = user.Age

	err = u.UserRepository.UpdateUser(userID, userData)
	if err != nil {
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
