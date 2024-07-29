package impl

import (
	"errors"

	"github.com/dieg0code/player-profile/src/auth"
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/repository"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	PasswordHasher services.PasswordHasher
	Validate       *validator.Validate
}

// Login implements services.AuthService.
func (a *AuthServiceImpl) Login(loginRequest request.LoginRequest) (*response.LoginResponse, error) {

	err := a.Validate.Struct(loginRequest)
	if err != nil {
		return nil, err
	}

	user, err := a.UserRepository.FindByEmail(loginRequest.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = a.PasswordHasher.ComparePassword(user.PassWord, loginRequest.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	loginResponse := &response.LoginResponse{
		Token: token,
	}

	return loginResponse, nil
}

func NewAuthService(userRepository repository.UserRepository, passwordHasher services.PasswordHasher, validate *validator.Validate) services.AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		PasswordHasher: passwordHasher,
		Validate:       validate,
	}
}
