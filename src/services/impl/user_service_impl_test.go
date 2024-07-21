package impl

import (
	"errors"
	"testing"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func TestUserServiceImpl_Create(t *testing.T) {
	t.Run("CreateUser_Scuccess", func(t *testing.T) {
		// Mocks
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		testUser := request.CreateUserRequest{
			UserName: "test",
			Password: "12345678",
			Email:    "test@test.com",
			Age:      18,
		}

		hashedPassword := "$2a$10$somehashedpassword"
		mockPasswordHasher.On("HashPassword", testUser.Password).Return(hashedPassword, nil)

		mockUserRepo.On("CreateUser", &models.User{
			UserName: testUser.UserName,
			PassWord: hashedPassword,
			Email:    testUser.Email,
			Age:      testUser.Age,
		}).Return(nil)

		err := userService.Create(testUser)

		require.NoError(t, err)
		mockUserRepo.AssertExpectations(t)

	})

	t.Run("CreateUser_Error", func(t *testing.T) {
		// Mocks
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		testUser := request.CreateUserRequest{
			UserName: "t",
			Password: "12",
			Email:    "asdasdasd",
			Age:      0,
		}

		err := userService.Create(testUser)

		require.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("CreateUser_HashPasswordError", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		testUser := request.CreateUserRequest{
			UserName: "test",
			Password: "12345678",
			Email:    "test@test.com",
			Age:      18,
		}

		mockPasswordHasher.On("HashPassword", testUser.Password).Return("", errors.New("hash error"))

		err := userService.Create(testUser)

		require.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("CreateUser_RepositoryError", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		testUser := request.CreateUserRequest{
			UserName: "test",
			Password: "12345678",
			Email:    "test@test.com",
			Age:      18,
		}

		hashedPassword := "$2a$10$somehashedpassword"
		mockPasswordHasher.On("HashPassword", testUser.Password).Return(hashedPassword, nil)

		mockUserRepo.On("CreateUser", &models.User{
			UserName: testUser.UserName,
			PassWord: hashedPassword,
			Email:    testUser.Email,
			Age:      testUser.Age,
		}).Return(errors.New("repository error"))

		err := userService.Create(testUser)

		require.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("CreateUser_ValidationError", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		invalidUser := request.CreateUserRequest{
			UserName: "test",
			Password: "short",
			Email:    "invalid-email",
			Age:      -1,
		}

		err := userService.Create(invalidUser)

		require.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}
