package impl

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/repository"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

// Create implements services.UserService.
func (u *UserServiceImpl) Create(user request.CreateUserRequest) error {

	err := u.Validate.Struct(user)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userModel := models.User{
		UserName: user.UserName,
		PassWord: string(hashedPassword),
		Email:    user.Email,
		Age:      user.Age,
	}

	err = u.UserRepository.CreateUser(&userModel)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements services.UserService.
func (u *UserServiceImpl) Delete(userID uint) error {

	err := u.UserRepository.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements services.UserService.
func (u *UserServiceImpl) GetAll() ([]response.UserResponse, error) {
	panic("unimplemented")
}

// GetByID implements services.UserService.
func (u *UserServiceImpl) GetByID(userID uint) (*response.UserResponse, error) {

	user, err := u.UserRepository.GetUser(userID)
	if err != nil {
		return nil, err
	}

	userResponse := response.UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Age:      user.Age,
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

	err = u.UserRepository.UpdateUser(userData)
	if err != nil {
		return err
	}

	return nil
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) services.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}
