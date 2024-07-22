package impl

import (
	"testing"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
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
