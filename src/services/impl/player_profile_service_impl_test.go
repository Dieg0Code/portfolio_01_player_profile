package impl

import (
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

func TestPlayerProfileServiceImpl_Create(t *testing.T) {
	t.Run("CreatePlayer_Success", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfile := request.CreatePlayerProfileRequest{
			Nickname:   "TestPlayer",
			Avatar:     "http://example.com/avatar.png",
			Level:      1,
			Experience: 10,
			Points:     5,
			UserID:     1,
		}

		// Expectations
		mockPlayerRepo.On("CreatePlayerProfile", &models.PlayerProfile{
			Nickname:   playerProfile.Nickname,
			Avatar:     playerProfile.Avatar,
			Level:      playerProfile.Level,
			Experience: playerProfile.Experience,
			Points:     playerProfile.Points,
			UserID:     playerProfile.UserID,
		}).Return(nil)

		// Execution
		err := playerService.Create(playerProfile)

		// Assertions
		require.NoError(t, err, "Error creating player profile")
		mockPlayerRepo.AssertExpectations(t)
	})

	t.Run("CreatePlayer_Error", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfile := request.CreatePlayerProfileRequest{
			Nickname: "TestPlayer",
			Avatar:   "http://example.com/avatar.png",
		}

		// Execution
		err := playerService.Create(playerProfile)

		// Assertions
		require.Error(t, err, "Expected error creating player profile with missing data")
		mockPlayerRepo.AssertExpectations(t)
	})

	t.Run("CreatePlayer_ValidationError", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfile := request.CreatePlayerProfileRequest{
			Nickname:   "TestPlayer",
			Avatar:     "http://example.com/avatar.png",
			Level:      0,
			Experience: 0,
			Points:     0,
			UserID:     0,
		}

		// Execution
		err := playerService.Create(playerProfile)

		// Assertions
		require.Error(t, err, "Expected error creating player with invalid data")
		require.EqualError(t, err, helpers.ErrPlayerProfileDataValidation.Error(), "Expected error creating player with invalid data")
		mockPlayerRepo.AssertExpectations(t)
	})

	t.Run("CreatePlayer_RepositoryError", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfile := request.CreatePlayerProfileRequest{
			Nickname:   "TestPlayer",
			Avatar:     "http://example.com/avatar.png",
			Level:      1,
			Experience: 10,
			Points:     5,
			UserID:     1,
		}

		// Expectations
		mockPlayerRepo.On("CreatePlayerProfile", &models.PlayerProfile{
			Nickname:   playerProfile.Nickname,
			Avatar:     playerProfile.Avatar,
			Level:      playerProfile.Level,
			Experience: playerProfile.Experience,
			Points:     playerProfile.Points,
			UserID:     playerProfile.UserID,
		}).Return(helpers.ErrRepository)

		// Execution
		err := playerService.Create(playerProfile)

		// Assertions
		require.Error(t, err, "Error creating player profile")
		require.Equal(t, helpers.ErrRepository, err, "Error creating player profile")
		mockPlayerRepo.AssertExpectations(t)
	})
}

func TestPlayerProfileServiceImpl_Delete(t *testing.T) {
	t.Run("DeletePlayer_Success", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfileID := uint(1)

		// Expectations
		mockPlayerRepo.On("DeletePlayerProfile", playerProfileID).Return(nil)

		// Execution
		err := playerService.Delete(playerProfileID)

		// Assertions
		require.NoError(t, err, "Error deleting player profile")
		mockPlayerRepo.AssertExpectations(t)
	})

	t.Run("DeletePlayer_InvalidID", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfileID := uint(0)

		// Execution
		err := playerService.Delete(playerProfileID)

		// Assertions
		require.Error(t, err, "Expected error deleting player profile with invalid ID")
		require.EqualError(t, err, helpers.ErrInvalidPlayerProfileID.Error(), "Expected error deleting player profile with invalid ID")
		mockPlayerRepo.AssertExpectations(t)
	})

	t.Run("DeletePlayer_RepositoryError", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfileID := uint(1)

		// Expectations
		mockPlayerRepo.On("DeletePlayerProfile", playerProfileID).Return(helpers.ErrRepository)

		// Execution
		err := playerService.Delete(playerProfileID)

		// Assertions
		require.Error(t, err, "Error deleting player profile")
		require.Equal(t, helpers.ErrRepository, err, "Error deleting player profile")
		mockPlayerRepo.AssertExpectations(t)
	})
}

func TestPlayerProfileServiceImpl_GetAll(t *testing.T) {}
func TestPlayerProfileServiceImpl_GetByID(t *testing.T) {
	t.Run("GetPlayer_Success", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfileID := uint(1)
		playerProfile := models.PlayerProfile{
			Model:      gorm.Model{ID: playerProfileID},
			Nickname:   "TestPlayer",
			Avatar:     "http://example.com/avatar.png",
			Level:      1,
			Experience: 10,
			Points:     5,
			UserID:     1,
		}

		// Expectations
		mockPlayerRepo.On("GetPlayerProfile", playerProfileID).Return(&playerProfile, nil)

		// Execution
		result, err := playerService.GetByID(playerProfileID)

		// Assertions
		require.NoError(t, err, "Error getting player profile")
		require.Equal(t, playerProfileID, result.ID, "Error getting player profile")
		mockPlayerRepo.AssertExpectations(t)
	})

	t.Run("GetPlayer_InvalidID", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfileID := uint(0)

		// Execution
		result, err := playerService.GetByID(playerProfileID)

		// Assertions
		require.Error(t, err, "Expected error getting player profile with invalid ID")
		require.Nil(t, result, "Expected nil result getting player profile with invalid ID")
		require.EqualError(t, err, helpers.ErrInvalidPlayerProfileID.Error(), "Expected error getting player profile with invalid ID")
		mockPlayerRepo.AssertExpectations(t)
	})

	t.Run("GetPlayer_RepositoryError", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfileID := uint(1)

		// Expectations
		mockPlayerRepo.On("GetPlayerProfile", playerProfileID).Return(nil, helpers.ErrRepository)

		// Execution
		result, err := playerService.GetByID(playerProfileID)

		// Assertions
		require.Error(t, err, "Error getting player profile")
		require.Nil(t, result, "Expected nil result getting player profile")
		require.Equal(t, helpers.ErrRepository, err, "Error getting player profile")
		mockPlayerRepo.AssertExpectations(t)
	})
}

func TestPlayerProfileServiceImpl_Update(t *testing.T) {
	t.Run("UpdatePlayer_Success", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Test data
		playerProfileID := uint(1)
		playerProfile := request.UpdatePlayerProfileRequest{
			Nickname:   "TestPlayer",
			Avatar:     "http://example.com/avatar.png",
			Level:      1,
			Experience: 10,
			Points:     5,
		}

		playerData := &models.PlayerProfile{
			Model:      gorm.Model{ID: playerProfileID},
			Nickname:   playerProfile.Nickname,
			Avatar:     playerProfile.Avatar,
			Level:      playerProfile.Level,
			Experience: playerProfile.Experience,
			Points:     playerProfile.Points,
		}

		// Expectations
		mockPlayerRepo.On("GetPlayerProfile", playerProfileID).Return(playerData, nil)
		mockPlayerRepo.On("UpdatePlayerProfile", playerProfileID, playerData).Return(nil)

		// Execution
		err := playerService.Update(playerProfileID, playerProfile)

		// Assertions
		require.NoError(t, err, "Error updating player profile")
		mockPlayerRepo.AssertExpectations(t)
	})

	t.Run("UpdatePlayer_InvalidID", func(t *testing.T) {
		// Mocks
		mockPlayerRepo := new(mocks.PlayerProfileRepository)
		mockValidator := validator.New()
		playerService := NewPlayerProfileServiceImpl(mockPlayerRepo, mockValidator)

		// Mock expectation for GetPlayerProfile
		mockPlayerRepo.On("GetPlayerProfile", mock.Anything).Return(nil, helpers.ErrorPlayerProfileNotFound)

		// Test data
		playerProfileID := uint(0)
		playerProfile := request.UpdatePlayerProfileRequest{
			Nickname:   "TestPlayer",
			Avatar:     "http://example.com/avatar.png",
			Level:      1,
			Experience: 10,
			Points:     5,
		}

		// Execution
		err := playerService.Update(playerProfileID, playerProfile)

		// Assertions
		require.Error(t, err, "Expected error updating player profile with invalid ID")
		require.EqualError(t, err, helpers.ErrorPlayerProfileNotFound.Error(), "Expected error updating player profile with invalid ID")
		mockPlayerRepo.AssertExpectations(t)
	})
}