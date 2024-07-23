package impl

import (
	"testing"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
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

func TestAchievementServiceImpl_Delete(t *testing.T) {
	t.Run("DeleteAchievement_Success", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(1)

		// Expectations
		mockAchievementRepo.On("DeleteAchievement", achievementID).Return(nil)

		// Execution
		err := achievementService.Delete(achievementID)

		// Assertions
		require.NoError(t, err, "Error deleting achievement")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("DeleteAchievement_InvalidID", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(0)

		// Execution
		err := achievementService.Delete(achievementID)

		// Assertions
		require.Error(t, err, "Expected error deleting achievement with invalid ID")
		require.EqualError(t, err, helpers.ErrInvalidAchievementID.Error(), "Expected error deleting achievement with invalid ID")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("DeleteAchievement_Error", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(1)

		// Expectations
		mockAchievementRepo.On("DeleteAchievement", achievementID).Return(helpers.ErrAchievementRepository)

		// Execution
		err := achievementService.Delete(achievementID)

		// Assertions
		require.Error(t, err, "Expected error deleting achievement")
		mockAchievementRepo.AssertExpectations(t)
	})
}

func TestAchievementServiceImpl_GetByID(t *testing.T) {
	t.Run("GetAchievement_Success", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(1)
		achievement := models.Achievement{
			Model:           gorm.Model{ID: achievementID},
			Name:            "Test",
			Description:     "Test Description",
			PlayerProfileID: 1,
		}

		// Expectations
		mockAchievementRepo.On("GetAchievement", achievementID).Return(&achievement, nil)

		// Execution
		result, err := achievementService.GetByID(achievementID)

		// Assertions
		require.NoError(t, err, "Error getting achievement")
		require.NotNil(t, result, "Expected achievement get nil")
		require.Equal(t, achievementID, result.ID, "Expected achievement ID")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("GetAchievement_InvalidID", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(0)

		// Execution
		result, err := achievementService.GetByID(achievementID)

		// Assertions
		require.Error(t, err, "Expected error getting achievement with invalid ID")
		require.Nil(t, result, "Expected achievement get nil")
		require.EqualError(t, err, helpers.ErrInvalidAchievementID.Error(), "Expected error getting achievement with invalid ID")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("GetAchievement_Error", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(1)

		// Expectations
		mockAchievementRepo.On("GetAchievement", achievementID).Return(nil, helpers.ErrAchievementRepository)

		// Execution
		result, err := achievementService.GetByID(achievementID)

		// Assertions
		require.Error(t, err, "Expected error getting achievement")
		require.Nil(t, result, "Expected achievement get nil")
		mockAchievementRepo.AssertExpectations(t)
	})
}

func TestAchievementServiceImpl_Update(t *testing.T) {
	t.Run("UpdateAchievement_Success", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(1)
		achievement := request.UpdateAchievementRequest{
			Name:        "Updated name",
			Description: "Updated Test Description",
		}

		// Expectations
		mockAchievementRepo.On("GetAchievement", achievementID).Return(&models.Achievement{}, nil)
		mockAchievementRepo.On("UpdateAchievement", achievementID, &models.Achievement{
			Name:        achievement.Name,
			Description: achievement.Description,
		}).Return(nil)

		// Execution
		err := achievementService.Update(achievementID, achievement)

		// Assertions
		require.NoError(t, err, "Error updating achievement")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("UpdateAchievement_InvalidID", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(0)
		achievement := request.UpdateAchievementRequest{
			Name:        "Updated name",
			Description: "Updated Test Description",
		}

		// Execution
		err := achievementService.Update(achievementID, achievement)

		// Assertions
		require.Error(t, err, "Expected error updating achievement with invalid ID")
		require.EqualError(t, err, helpers.ErrInvalidAchievementID.Error(), "Expected error updating achievement with invalid ID")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("UpdateAchievement_Error", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(1)
		achievement := request.UpdateAchievementRequest{
			Name:        "Updated name",
			Description: "Updated Test Description",
		}

		// Expectations
		mockAchievementRepo.On("GetAchievement", achievementID).Return(nil, helpers.ErrAchievementRepository)

		// Execution
		err := achievementService.Update(achievementID, achievement)

		// Assertions
		require.Error(t, err, "Expected error updating achievement")
		mockAchievementRepo.AssertExpectations(t)
	})
}
