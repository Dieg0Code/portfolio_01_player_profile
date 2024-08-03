package impl

import (
	"errors"
	"strings"
	"testing"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAuthServiceImpl(t *testing.T) {
	t.Run("Login_Success", func(t *testing.T) {
		// Mocks
		mockUserRepo := new(mocks.UserRepository)
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		validate := validator.New()
		mockAuthUtils := new(mocks.MockAuthUtils)
		authService := NewAuthService(mockUserRepo, mockPasswordHasher, validate, mockAuthUtils)

		// Test data
		loginRequest := request.LoginRequest{
			Email:    "test@test.com",
			Password: "password",
		}

		mockUserRepo.On("FindByEmail", loginRequest.Email).Return(&models.User{
			Model:    gorm.Model{ID: 1},
			UserName: "test",
			Email:    "test@test.com",
			PassWord: "password",
			Age:      20,
			Role:     "admin",
		}, nil)

		mockPasswordHasher.On("ComparePassword", "password", loginRequest.Password).Return(nil)

		mockAuthUtils.On("GenerateToken", uint(1), "admin").Return("token", nil)

		// Execution
		loginResponse, err := authService.Login(loginRequest)

		// Validation
		assert.Nil(t, err)
		assert.NotNil(t, loginResponse)
		assert.Equal(t, "token", loginResponse.Token)

		mockUserRepo.AssertExpectations(t)
		mockPasswordHasher.AssertExpectations(t)
		mockAuthUtils.AssertExpectations(t)

	})

	t.Run("Login_Fail_InvalidRequest", func(t *testing.T) {
		// Mocks
		mockUserRepo := new(mocks.UserRepository)
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		validate := validator.New()
		mockAuthUtils := new(mocks.MockAuthUtils)
		authService := NewAuthService(mockUserRepo, mockPasswordHasher, validate, mockAuthUtils)

		// Test data
		loginRequest := request.LoginRequest{
			Email:    "",
			Password: "",
		}

		// Execution
		loginResponse, err := authService.Login(loginRequest)

		// Validation
		assert.NotNil(t, err)
		assert.Nil(t, loginResponse)

		mockUserRepo.AssertExpectations(t)
		mockPasswordHasher.AssertExpectations(t)
		mockAuthUtils.AssertExpectations(t)
	})

	t.Run("Login_Fail_UserNotFound", func(t *testing.T) {
		// Mocks
		mockUserRepo := new(mocks.UserRepository)
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		validate := validator.New()
		mockAuthUtils := new(mocks.MockAuthUtils)
		authService := NewAuthService(mockUserRepo, mockPasswordHasher, validate, mockAuthUtils)

		// Test data
		loginRequest := request.LoginRequest{
			Email:    "nonexistent@test.com",
			Password: "password",
		}

		mockUserRepo.On("FindByEmail", loginRequest.Email).Return(nil, assert.AnError)

		// Execution
		loginResponse, err := authService.Login(loginRequest)

		// Validation
		assert.NotNil(t, err, "Expected error finding nonexistent user")
		assert.Nil(t, loginResponse, "Expected nil in login response")

		mockUserRepo.AssertExpectations(t)
		mockPasswordHasher.AssertExpectations(t)
		mockAuthUtils.AssertExpectations(t)
	})

	t.Run("Login_Fail_InvalidPassword", func(t *testing.T) {
		// Mocks
		mockUserRepo := new(mocks.UserRepository)
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		validate := validator.New()
		mockAuthUtils := new(mocks.MockAuthUtils)
		authService := NewAuthService(mockUserRepo, mockPasswordHasher, validate, mockAuthUtils)

		// Test data
		loginRequest := request.LoginRequest{
			Email:    "test@test.com",
			Password: "wrongpassword",
		}

		mockUserRepo.On("FindByEmail", loginRequest.Email).Return(&models.User{
			Model:    gorm.Model{ID: 1},
			UserName: "test",
			Email:    "test@test.com",
			PassWord: "correctpassword",
			Age:      20,
			Role:     "admin",
		}, nil)

		mockPasswordHasher.On("ComparePassword", "correctpassword", loginRequest.Password).Return(errors.New("password mismatch"))

		// Execution
		loginResponse, err := authService.Login(loginRequest)

		// Validation
		assert.NotNil(t, err, "Expected error for invalid password")
		assert.Nil(t, loginResponse, "Expected nil in login response")

		mockUserRepo.AssertExpectations(t)
		mockPasswordHasher.AssertExpectations(t)
		mockAuthUtils.AssertExpectations(t)
	})

	t.Run("Login_Fail_TokenGeneration", func(t *testing.T) {
		// Mocks
		mockUserRepo := new(mocks.UserRepository)
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		validate := validator.New()
		mockAuthUtils := new(mocks.MockAuthUtils)
		authService := NewAuthService(mockUserRepo, mockPasswordHasher, validate, mockAuthUtils)

		// Test data
		loginRequest := request.LoginRequest{
			Email:    "test@test.com",
			Password: "password",
		}

		mockUserRepo.On("FindByEmail", loginRequest.Email).Return(&models.User{
			Model:    gorm.Model{ID: 1},
			UserName: "test",
			Email:    "test@test.com",
			PassWord: "password",
			Age:      20,
			Role:     "admin",
		}, nil)

		mockPasswordHasher.On("ComparePassword", "password", loginRequest.Password).Return(nil)

		mockAuthUtils.On("GenerateToken", uint(1), "admin").Return("", errors.New("token generation failed"))

		// Execution
		loginResponse, err := authService.Login(loginRequest)

		// Validation
		assert.NotNil(t, err, "Expected error for token generation failure")
		assert.Nil(t, loginResponse, "Expected nil in login response")

		mockUserRepo.AssertExpectations(t)
		mockPasswordHasher.AssertExpectations(t)
		mockAuthUtils.AssertExpectations(t)
	})

	t.Run("Login_Fail_LongPassword", func(t *testing.T) {
		// Mocks
		mockUserRepo := new(mocks.UserRepository)
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		validate := validator.New()
		mockAuthUtils := new(mocks.MockAuthUtils)
		authService := NewAuthService(mockUserRepo, mockPasswordHasher, validate, mockAuthUtils)

		// Test data
		longPassword := "a" + strings.Repeat("b", 4096) // assuming a password length limit
		loginRequest := request.LoginRequest{
			Email:    "test@test.com",
			Password: longPassword,
		}

		mockUserRepo.On("FindByEmail", loginRequest.Email).Return(nil, assert.AnError)

		// Execution
		loginResponse, err := authService.Login(loginRequest)

		// Validation
		assert.NotNil(t, err, "Expected validation error for long password")
		assert.Nil(t, loginResponse, "Expected nil in login response")

		mockUserRepo.AssertExpectations(t)
		mockPasswordHasher.AssertExpectations(t)
		mockAuthUtils.AssertExpectations(t)
	})

}
