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
			Name:        "Test Achievement",
			Description: "Test Description",
		}

		// Expectations
		mockAchievementRepo.On("CreateAchievement", &models.Achievement{
			Name:        achievement.Name,
			Description: achievement.Description,
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
			Name:        "",
			Description: "",
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
			Model:       gorm.Model{ID: achievementID},
			Name:        "Test",
			Description: "Test Description",
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

func TestAchievementServiceImpl_GetAll(t *testing.T) {
	t.Run("GetAllAchievements_Success", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievement := models.Achievement{
			Model:       gorm.Model{ID: 1},
			Name:        "Test",
			Description: "Test Description",
		}

		achivement1 := models.Achievement{
			Model:       gorm.Model{ID: 2},
			Name:        "Test 2",
			Description: "Test Description 2",
		}

		var respnseMock []models.Achievement
		respnseMock = append(respnseMock, achievement, achivement1)

		// Expectations
		mockAchievementRepo.On("GetAllAchievements", 0, 10).Return(respnseMock, nil)

		// Execution
		result, err := achievementService.GetAll(1, 10)

		// Assertions
		require.NoError(t, err, "Error getting all achievements")
		require.NotNil(t, result, "Expected achievements get nil")
		require.Len(t, result, 2, "Expected achievements length")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("GetAllAchievements_Error", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Expectations
		mockAchievementRepo.On("GetAllAchievements", 0, 10).Return([]models.Achievement{}, helpers.ErrAchievementRepository)

		// Execution
		result, err := achievementService.GetAll(1, 10)

		// Assertions
		require.Error(t, err, "Expected error getting all achievements")
		require.Nil(t, result, "Expected achievements get nil")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("GetAllAchievements_Empty", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Expectations
		mockAchievementRepo.On("GetAllAchievements", 0, 10).Return([]models.Achievement{}, nil)

		// Execution
		result, err := achievementService.GetAll(1, 10)

		// Assertions
		require.NoError(t, err, "Error getting all achievements")
		require.Empty(t, result, "Expected achievements get empty")
		require.Len(t, result, 0, "Expected achievements length")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("GetAllAchievements_RepositoryError", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Expectations
		mockAchievementRepo.On("GetAllAchievements", 0, 10).Return([]models.Achievement{}, helpers.ErrAchievementRepository)

		// Execution
		result, err := achievementService.GetAll(1, 10)

		// Assertions
		require.Error(t, err, "Expected error getting all achievements")
		require.Nil(t, result, "Expected achievements get nil")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("GetAllAchievements_InvalidPagination", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Execution
		result, err := achievementService.GetAll(0, 0)

		// Assertions
		require.Error(t, err, "Expected error getting all achievements")
		require.Nil(t, result, "Expected achievements get nil")
		require.EqualError(t, err, helpers.ErrInvalidPagination.Error(), "Expected error getting all achievements")
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

func TetAchievementServiceImpl_GetAchievementWithPlayers(t *testing.T) {
	t.Run("GetAchievementWithPlayers_Success", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(1)
		achievement := models.Achievement{
			Model:       gorm.Model{ID: achievementID},
			Name:        "Test",
			Description: "Test Description",
			PlayerProfiles: []models.PlayerProfile{
				{
					Model:      gorm.Model{ID: 1},
					Nickname:   "test nickaname",
					Avatar:     "test.png",
					Level:      12,
					Experience: 10,
					Points:     100,
				},
			},
		}

		// Expectations
		mockAchievementRepo.On("GetAchievementWithPlayers", achievementID).Return(&achievement, nil)

		// Execution
		result, err := achievementService.GetAchievementWithPlayers(achievementID)

		// Assertions
		require.NoError(t, err, "Error getting achievement with players")
		require.NotNil(t, result, "Expected achievement get nil")
		require.Equal(t, achievementID, result.ID, "Expected achievement ID")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("GetAchievementWithPlayers_InvalidID", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(0)

		mockAchievementRepo.On("GetAchievementWithPlayers", achievementID).Return(nil, helpers.ErrInvalidAchievementID)

		// Execution
		result, err := achievementService.GetAchievementWithPlayers(achievementID)

		// Assertions
		require.Error(t, err, "Expected error getting achievement with players with invalid ID")
		require.Nil(t, result, "Expected achievement get nil")
		require.EqualError(t, err, helpers.ErrInvalidAchievementID.Error(), "Expected error getting achievement with players with invalid ID")
		mockAchievementRepo.AssertExpectations(t)
	})

	t.Run("GetAchievementWithPlayers_RepositoyError", func(t *testing.T) {
		// Mocks
		mockAchievementRepo := new(mocks.AchievementRepository)
		mockValidator := validator.New()
		achievementService := NewAchievementServiceImpl(mockAchievementRepo, mockValidator)

		// Test data
		achievementID := uint(1)

		mockAchievementRepo.On("GetAchievementWithPlayers", achievementID).Return(nil, helpers.ErrAchievementRepository)

		// Execution
		result, err := achievementService.GetAchievementWithPlayers(achievementID)

		// Assertions
		require.Error(t, err, "Expected error getting achievement with players")
		require.Nil(t, result, "Expected achievement get nil")
		mockAchievementRepo.AssertExpectations(t)
	})
}
