package impl

import (
	"testing"

	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils"
	"github.com/stretchr/testify/require"
)

func TestAchievementRepository_CreateAchievement(t *testing.T) {

	t.Run("CreateAchievement_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		userRepo := NewUserRepositoryImpl(db)
		playerProfileRepo := NewPlayerProfileRepositoryImpl(db)
		achievementRepo := NewAchievementRepositoryImpl(db)

		// Create User
		user := &models.User{
			UserName: "testuser",
			PassWord: "testpass",
			Email:    "test@example.com",
			Age:      25,
		}
		err := userRepo.CreateUser(user)
		require.NoError(t, err, "Error creating user")

		// Create PlayerProfile
		playerProfile := &models.PlayerProfile{
			Nickname:   "testnick",
			Avatar:     "test.png",
			Level:      1,
			Experience: 100,
			Points:     50,
			UserID:     user.ID,
		}
		err = playerProfileRepo.CreatePlayerProfile(playerProfile)
		require.NoError(t, err, "Error creating player profile")

		// Create Achievement
		achievement := &models.Achievement{
			Name:            "Test Achievement",
			Description:     "This is a test achievement",
			PlayerProfileID: playerProfile.ID,
		}
		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err, "Error creating achievement")

		// Verify Achievement
		createdAchievement, err := achievementRepo.GetAchievement(achievement.ID)
		require.NoError(t, err, "Error getting achievement")
		require.Equal(t, achievement.Name, createdAchievement.Name, "Achievement names do not match")
		require.Equal(t, achievement.Description, createdAchievement.Description, "Achievement descriptions do not match")
		require.Equal(t, achievement.PlayerProfileID, createdAchievement.PlayerProfileID, "Achievement player profile IDs do not match")
	})

	t.Run("CreateAchievement_Duplicate", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		testUser1 := &models.User{
			UserName: "test",
			PassWord: "test",
			Email:    "test@test.com",
			Age:      20,
		}

		userRepo := NewUserRepositoryImpl(db)
		playerProfileRepo := NewPlayerProfileRepositoryImpl(db)
		achievementRepo := NewAchievementRepositoryImpl(db)

		// Create user
		err := userRepo.CreateUser(testUser1)
		require.NoError(t, err, "Error creating user")

		testPlayerProfile1 := &models.PlayerProfile{
			Nickname:   "test",
			Avatar:     "test.png",
			Level:      1,
			Experience: 10,
			Points:     100,
			UserID:     testUser1.ID,
		}

		// Create player profile
		err = playerProfileRepo.CreatePlayerProfile(testPlayerProfile1)
		require.NoError(t, err, "Error creating player profile")

		testAchievement1 := &models.Achievement{
			Name:            "test",
			Description:     "test",
			PlayerProfileID: testPlayerProfile1.ID,
		}

		// Create achievement
		err = achievementRepo.CreateAchievement(testAchievement1)
		require.NoError(t, err, "Error creating achievement")

		// Attempt to create the same achievement
		err = achievementRepo.CreateAchievement(testAchievement1)
		require.Error(t, err, "Expected error creating duplicate achievement")
	})

}

func TestAchievementRepository_GetAchievement(t *testing.T) {

	t.Run("GetAchievement_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		userRepo := NewUserRepositoryImpl(db)
		playerProfileRepo := NewPlayerProfileRepositoryImpl(db)
		achievementRepo := NewAchievementRepositoryImpl(db)

		// Create User
		user := &models.User{
			UserName: "testuser",
			PassWord: "testpass",
			Email:    "test@test.com",
			Age:      25,
		}

		err := userRepo.CreateUser(user)
		require.NoError(t, err, "Error creating user")

		// Create PlayerProfile
		playerProfile := &models.PlayerProfile{
			Nickname:   "testnick",
			Avatar:     "test.png",
			Level:      1,
			Experience: 100,
			Points:     50,
			UserID:     user.ID,
		}

		err = playerProfileRepo.CreatePlayerProfile(playerProfile)
		require.NoError(t, err, "Error creating player profile")

		// Create Achievement
		achievement := &models.Achievement{
			Name:            "Test Achievement",
			Description:     "This is a test achievement",
			PlayerProfileID: playerProfile.ID,
		}

		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err, "Error creating achievement")

		// Verify Achievement
		createdAchievement, err := achievementRepo.GetAchievement(achievement.ID)
		require.NoError(t, err, "Error getting achievement")
		require.Equal(t, achievement.Name, createdAchievement.Name, "Achievement names do not match")
		require.Equal(t, achievement.Description, createdAchievement.Description, "Achievement descriptions do not match")
		require.Equal(t, achievement.PlayerProfileID, createdAchievement.PlayerProfileID, "Achievement player profile IDs do not match")
	})

	t.Run("GetAchievement_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		achievementRepo := NewAchievementRepositoryImpl(db)

		// Attempt to get non-existent achievement
		_, err := achievementRepo.GetAchievement(1)
		require.Error(t, err, "Expected error getting non-existent achievement")
	})
}

func TestAchievementRepository_UpdateAchievement(t *testing.T) {
	t.Run("UpdateAchievement_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		userRepo := NewUserRepositoryImpl(db)
		playerProfileRepo := NewPlayerProfileRepositoryImpl(db)
		achievementRepo := NewAchievementRepositoryImpl(db)

		// Create User
		user := &models.User{
			UserName: "testuser",
			PassWord: "testpass",
			Email:    "test@test.com",
			Age:      25,
		}

		err := userRepo.CreateUser(user)
		require.NoError(t, err, "Error creating user")

		// Create PlayerProfile
		playerProfile := &models.PlayerProfile{
			Nickname:   "testnick",
			Avatar:     "test.png",
			Level:      1,
			Experience: 100,
			Points:     50,
			UserID:     user.ID,
		}

		err = playerProfileRepo.CreatePlayerProfile(playerProfile)
		require.NoError(t, err, "Error creating player profile")

		// Create Achievement
		// Create Achievement
		achievement := &models.Achievement{
			Name:            "Test Achievement",
			Description:     "This is a test achievement",
			PlayerProfileID: playerProfile.ID,
		}

		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err, "Error creating achievement")

		// Update Achievement
		achievement.Name = "Updated Achievement"
		achievement.Description = "This is an updated test achievement"

		err = achievementRepo.UpdateAchievement(achievement)
		require.NoError(t, err, "Error updating achievement")

		// Verify Achievement
		updatedAchievement, err := achievementRepo.GetAchievement(achievement.ID)
		require.NoError(t, err, "Error getting achievement")
		require.Equal(t, achievement.Name, updatedAchievement.Name, "Achievement names do not match")
	})

	t.Run("UpdateAchievement_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		achievementRepo := NewAchievementRepositoryImpl(db)

		// Attempt to update non-existent achievement
		err := achievementRepo.UpdateAchievement(&models.Achievement{})
		require.Error(t, err, "Expected error updating non-existent achievement")
	})
}

func TestAchievementRepository_DeleteAchievement(t *testing.T) {
	t.Run("DeleteAchievement_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		userRepo := NewUserRepositoryImpl(db)
		playerProfileRepo := NewPlayerProfileRepositoryImpl(db)
		achievementRepo := NewAchievementRepositoryImpl(db)

		// Create User
		user := &models.User{
			UserName: "testuser",
			PassWord: "testpass",
			Email:    "test@test.com",
			Age:      25,
		}

		err := userRepo.CreateUser(user)
		require.NoError(t, err, "Error creating user")

		// Create PlayerProfile
		playerProfile := &models.PlayerProfile{
			Nickname:   "testnick",
			Avatar:     "test.png",
			Level:      1,
			Experience: 100,
			Points:     50,
			UserID:     user.ID,
		}

		err = playerProfileRepo.CreatePlayerProfile(playerProfile)
		require.NoError(t, err, "Error creating player profile")

		// Create Achievement
		achievement := &models.Achievement{
			Name:            "Test Achievement",
			Description:     "This is a test achievement",
			PlayerProfileID: playerProfile.ID,
		}

		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err, "Error creating achievement")

		// Delete Achievement
		err = achievementRepo.DeleteAchievement(achievement.ID)
		require.NoError(t, err, "Error deleting achievement")

		// Verify Achievement
		_, err = achievementRepo.GetAchievement(achievement.ID)
		require.Error(t, err, "Expected error getting deleted achievement")
	})

	t.Run("DeleteAchievement_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		achievementRepo := NewAchievementRepositoryImpl(db)

		// Attempt to delete non-existent achievement
		err := achievementRepo.DeleteAchievement(1)
		require.Error(t, err, "Expected error deleting non-existent achievement")
	})
}

func TestAchievementRepository_CheckAchievementExists(t *testing.T) {
	t.Run("CheckAchievementExists_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		userRepo := NewUserRepositoryImpl(db)
		playerProfileRepo := NewPlayerProfileRepositoryImpl(db)
		achievementRepo := NewAchievementRepositoryImpl(db)

		// Create User
		user := &models.User{
			UserName: "testuser",
			PassWord: "testpass",
			Email:    "test@test.com",
			Age:      25,
		}

		err := userRepo.CreateUser(user)
		require.NoError(t, err, "Error creating user")

		// Create PlayerProfile
		playerProfile := &models.PlayerProfile{
			Nickname:   "testnick",
			Avatar:     "test.png",
			Level:      1,
			Experience: 100,
			Points:     50,
			UserID:     user.ID,
		}

		err = playerProfileRepo.CreatePlayerProfile(playerProfile)
		require.NoError(t, err, "Error creating player profile")

		// Create Achievement
		achievement := &models.Achievement{
			Name:            "Test Achievement",
			Description:     "This is a test achievement",
			PlayerProfileID: playerProfile.ID,
		}

		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err, "Error creating achievement")

		// Check Achievement Exists
		exists, err := achievementRepo.CheckAchievementExists(achievement.ID)
		require.NoError(t, err, "Error checking achievement exists")
		require.True(t, exists, "Expected achievement to exist")
	})

	t.Run("CheckAchievementExists_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		achievementRepo := NewAchievementRepositoryImpl(db)

		// Check non-existent achievement
		exists, err := achievementRepo.CheckAchievementExists(1)
		require.NoError(t, err, "Error checking achievement exists")
		require.False(t, exists, "Expected achievement to not exist")
	})
}