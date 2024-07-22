package impl

import (
	"errors"
	"testing"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
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

func TestUserServiceImpl_Delete(t *testing.T) {
	t.Run("DeleteUser_Success", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		userID := uint(1)

		mockUserRepo.On("DeleteUser", userID).Return(nil)

		err := userService.Delete(userID)

		require.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("DeleteUser_Error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		userID := uint(1)

		mockUserRepo.On("DeleteUser", userID).Return(errors.New("repository error"))

		err := userService.Delete(userID)

		require.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("DeleteUser_ValidationError", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		err := userService.Delete(0)

		require.Error(t, err)
		require.Equal(t, err, helpers.ErrInvalidUserID)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserServiceImpl_GetById(t *testing.T) {
	t.Run("GetUser_Success", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		userID := uint(1)
		user := &models.User{
			Model:    gorm.Model{ID: userID},
			UserName: "test",
			PassWord: "hashedpassword",
			Email:    "test@test.com",
			Age:      18,
		}

		mockUserRepo.On("GetUser", userID).Return(user, nil)

		userResponse, err := userService.GetByID(userID)

		require.NoError(t, err)
		require.Equal(t, user.ID, userResponse.ID)
		require.Equal(t, user.UserName, userResponse.UserName)
	})

	t.Run("GetUser_Error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)

		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		userID := uint(1)
		mockUserRepo.On("GetUser", userID).Return(nil, errors.New("repository error"))

		userResponse, err := userService.GetByID(userID)

		require.Error(t, err)
		require.Nil(t, userResponse)
		require.Equal(t, "repository error", err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("GetUser_NotFound", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)

		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		userID := uint(1)

		mockUserRepo.On("GetUser", userID).Return(nil, nil)

		userResponse, err := userService.GetByID(userID)

		require.Error(t, err)
		require.Nil(t, userResponse)
		require.Equal(t, "user not found", err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("GetUser_ValidationError", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		mockPasswordHasher := new(mocks.MockPasswordHasher)

		userService := NewUserServiceImpl(mockUserRepo, mockValidator, mockPasswordHasher)

		userID := uint(1)
		mockUser := &models.User{
			Model:    gorm.Model{ID: userID},
			UserName: "t", // This should fail validation (too short)
			Email:    "invalid-email",
			Age:      17, // This should fail validation (under 18)
		}
		mockUserRepo.On("GetUser", userID).Return(mockUser, nil)

		userResponse, err := userService.GetByID(userID)

		require.Error(t, err)
		require.Nil(t, userResponse)
		require.ErrorIs(t, err, helpers.ErrUserDataValidation)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserServiceImpl_Update(t *testing.T) {
	t.Run("UpdateUser_Success", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, nil) // Assuming no password hasher needed for update

		userID := uint(1)
		updateRequest := request.UpdateUserRequest{
			UserName: "updatedName",
			Email:    "update@test.com",
			Age:      25,
		}

		existingUser := &models.User{
			Model:    gorm.Model{ID: userID},
			UserName: "originalName",
			Email:    "original@test.com",
			Age:      20,
		}

		mockUserRepo.On("GetUser", userID).Return(existingUser, nil)
		mockUserRepo.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(nil)

		err := userService.Update(userID, updateRequest)

		require.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("UpdateUser_GetUserError", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, nil)

		userID := uint(1)
		updateRequest := request.UpdateUserRequest{}

		mockUserRepo.On("GetUser", userID).Return(nil, errors.New("user not found"))

		err := userService.Update(userID, updateRequest)

		require.Error(t, err)
		require.Equal(t, "user not found", err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("UpdateUser_ValidationError", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, nil)

		userID := uint(1)
		updateRequest := request.UpdateUserRequest{
			UserName: "",
			Email:    "invalid-email",
			Age:      -1,
		}

		mockUserRepo.On("GetUser", userID).Return(&models.User{}, nil)

		err := userService.Update(userID, updateRequest)

		require.Error(t, err)
		mockUserRepo.AssertNotCalled(t, "UpdateUser", mock.AnythingOfType("*models.User"))
	})

	t.Run("UpdateUser_UpdateUserError", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockValidator := validator.New()
		userService := NewUserServiceImpl(mockUserRepo, mockValidator, nil)

		userID := uint(1)
		updateRequest := request.UpdateUserRequest{
			UserName: "updatedName",
			Email:    "update@test.com",
			Age:      25,
		}

		existingUser := &models.User{
			Model:    gorm.Model{ID: userID},
			UserName: "originalName",
			Email:    "original@test.com",
			Age:      20,
		}

		mockUserRepo.On("GetUser", userID).Return(existingUser, nil)
		mockUserRepo.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(errors.New("update error"))

		err := userService.Update(userID, updateRequest)

		require.Error(t, err)
		require.Equal(t, "update error", err.Error())
		mockUserRepo.AssertExpectations(t)
	})
}
