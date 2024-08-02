package impl

import (
	"testing"

	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
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
			Name:        "Test Achievement",
			Description: "This is a test achievement",
		}
		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err, "Error creating achievement")

		// Verify Achievement
		createdAchievement, err := achievementRepo.GetAchievement(achievement.ID)
		require.NoError(t, err, "Error getting achievement")
		require.Equal(t, achievement.Name, createdAchievement.Name, "Achievement names do not match")
		require.Equal(t, achievement.Description, createdAchievement.Description, "Achievement descriptions do not match")
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
			Name:        "test",
			Description: "test",
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
			Name:        "Test Achievement",
			Description: "This is a test achievement",
		}

		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err, "Error creating achievement")

		// Verify Achievement
		createdAchievement, err := achievementRepo.GetAchievement(achievement.ID)
		require.NoError(t, err, "Error getting achievement")
		require.Equal(t, achievement.Name, createdAchievement.Name, "Achievement names do not match")
		require.Equal(t, achievement.Description, createdAchievement.Description, "Achievement descriptions do not match")
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

func TwardAchievementRepository_GetAllAchievements(t *testing.T) {
	t.Run("GetAllAchievements_Success", func(t *testing.T) {
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
			Email:    "test1@test.com",
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

		// Create Achievements
		achievements1 := &models.Achievement{
			Name:        "Test Achievement 1",
			Description: "This is a test achievement 1",
		}

		achievements2 := &models.Achievement{
			Name:        "Test Achievement 2",
			Description: "This is a test achievement 2",
		}

		err = achievementRepo.CreateAchievement(achievements1)
		require.NoError(t, err, "Error creating achievement 1")

		err = achievementRepo.CreateAchievement(achievements2)
		require.NoError(t, err, "Error creating achievement 2")

		// Get all achievements
		allAchievements, err := achievementRepo.GetAllAchievements(0, 10)
		require.NoError(t, err, "Error getting all achievements")
		require.Len(t, allAchievements, 2, "Expected 2 achievements")
		require.Equal(t, achievements1.Name, allAchievements[0].Name, "Achievement 1 names do not match")
		require.Equal(t, achievements2.Name, allAchievements[1].Name, "Achievement 2 names do not match")
	})

	t.Run("GetAllAchievements_Empty", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		achievementRepo := NewAchievementRepositoryImpl(db)

		// Get all achievements when there are none
		allAchievements, err := achievementRepo.GetAllAchievements(0, 10)
		require.NoError(t, err, "Error getting all achievements")
		require.Len(t, allAchievements, 0, "Expected 0 achievements")
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
		achievementID := uint(1)
		achievement := &models.Achievement{
			Model:       gorm.Model{ID: achievementID},
			Name:        "Test Achievement",
			Description: "This is a test achievement",
		}

		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err, "Error creating achievement")

		// Update Achievement
		achievement.Name = "Updated Achievement"
		achievement.Description = "This is an updated test achievement"

		err = achievementRepo.UpdateAchievement(achievement.ID, achievement)
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

		achievementID := uint(0)

		// Attempt to update non-existent achievement
		err := achievementRepo.UpdateAchievement(achievementID, &models.Achievement{})
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
			Name:        "Test Achievement",
			Description: "This is a test achievement",
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
			Name:        "Test Achievement",
			Description: "This is a test achievement",
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

func TestAchievementRepository_GetAchievementWithPlayers(t *testing.T) {
	t.Run("GetAchievementWithPlayers_Success", func(t *testing.T) {
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
			Email:    "test@test",
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
			Name:        "Test Achievement",
			Description: "This is a test achievement",
			PlayerProfiles: []models.PlayerProfile{
				*playerProfile,
			},
		}

		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err, "Error creating achievement")

		// Get Achievement with Players
		achievementWithPlayers, err := achievementRepo.GetAchievementWithPlayers(achievement.ID)
		require.NoError(t, err, "Error getting achievement with players")
		require.Equal(t, achievement.Name, achievementWithPlayers.Name, "Achievement names do not match")
		require.Equal(t, achievement.Description, achievementWithPlayers.Description, "Achievement descriptions do not match")
		require.Len(t, achievementWithPlayers.PlayerProfiles, 1, "Expected 1 player profile")
		require.Equal(t, playerProfile.Nickname, achievementWithPlayers.PlayerProfiles[0].Nickname, "Player profile nicknames do not match")
	})

	t.Run("GetAchievementWithPlayers_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()
		achievementRepo := NewAchievementRepositoryImpl(db)

		// Get non-existent achievement with players
		_, err := achievementRepo.GetAchievementWithPlayers(1)
		require.Error(t, err, "Expected error getting achievement with players")
	})
}
