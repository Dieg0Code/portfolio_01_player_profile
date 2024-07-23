package impl

import (
	"testing"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func TestAchievementServiceImpl_Create(t *testing.T) {

	t.Run("CreateAchievement_Success", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievement := request.CreateAchievementRequest{
			Name:            "Test Achievement",
			Description:     "Test Description",
			PlayerProfileID: 1,
		}

		// Expectations
		mockAchievementRepo.On("CreateAchievement", &models.Achievement{
			Name:            achievement.Name,
			Description:     achievement.Description,
			PlayerProfileID: achievement.PlayerProfileID,
		}).Return(nil)

		// Execution
		err := achievementService.Create(achievement)

		// Assertions
		require.NoError(t, err, "Error creating achievement")
		mockAchievementRepo.AssertExpectations(t)

	})

	t.Run("CreateAchievement_ValidationError", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievement := request.CreateAchievementRequest{
			Name:            "",
			Description:     "",
			PlayerProfileID: 0,
		}

		// Execution
		err := achievementService.Create(achievement)

		// Assertions
		require.Error(t, err, "Error creating achievement")
		mockAchievementRepo.AssertExpectations(t)
	})
}
