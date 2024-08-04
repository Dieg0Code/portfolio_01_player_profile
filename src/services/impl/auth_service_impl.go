package impl

import (
	"errors"

	"github.com/dieg0code/player-profile/src/auth"
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/repository"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	PasswordHasher services.PasswordHasher
	Validate       *validator.Validate
	AuthUtils      auth.AuthUtils
}

// Login implements services.AuthService.
func (a *AuthServiceImpl) Login(loginRequest request.LoginRequest) (*response.LoginResponse, error) {
	// Validación de la solicitud
	err := a.Validate.Struct(loginRequest)
	if err != nil {
		logrus.WithError(err).Error("[AuthServiceImpl.Login] Failed to validate login request")
		return nil, err
	}

	// Búsqueda del usuario por correo electrónico
	user, err := a.UserRepository.FindByEmail(loginRequest.Email)
	if err != nil {
		logrus.WithError(err).Error("[AuthServiceImpl.Login] Failed to find user by email")
		return nil, errors.New("user not found")
	}

	// Comparación de la contraseña
	err = a.PasswordHasher.ComparePassword(user.PassWord, loginRequest.Password)
	if err != nil {
		logrus.WithError(err).Error("[AuthServiceImpl.Login] Failed to compare password")
		return nil, errors.New("invalid credentials")
	}

	// Generación del token
	token, err := a.AuthUtils.GenerateToken(user.ID, user.Role)
	if err != nil {
		logrus.WithError(err).Error("[AuthServiceImpl.Login] Failed to generate token")
		return nil, errors.New("failed to generate token")
	}

	// Creación de la respuesta de inicio de sesión
	loginResponse := &response.LoginResponse{
		Token: token,
	}

	return loginResponse, nil
}

func NewAuthService(userRepository repository.UserRepository, passwordHasher services.PasswordHasher, validate *validator.Validate, auth auth.AuthUtils) services.AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		PasswordHasher: passwordHasher,
		Validate:       validate,
		AuthUtils:      auth,
	}
}
