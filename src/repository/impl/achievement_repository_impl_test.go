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
		require.NoError(t, err)

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
		require.NoError(t, err)

		// Create Achievement
		achievement := &models.Achievement{
			Name:            "Test Achievement",
			Description:     "This is a test achievement",
			PlayerProfileID: playerProfile.ID,
		}
		err = achievementRepo.CreateAchievement(achievement)
		require.NoError(t, err)

		// Verify Achievement
		createdAchievement, err := achievementRepo.GetAchievement(int(achievement.ID))
		require.NoError(t, err)
		require.Equal(t, achievement.Name, createdAchievement.Name)
		require.Equal(t, achievement.Description, createdAchievement.Description)
		require.Equal(t, achievement.PlayerProfileID, createdAchievement.PlayerProfileID)
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
