package impl

import (
	"testing"

	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils"
	"github.com/stretchr/testify/require"
)

var (
	testUser = &models.User{
		UserName: "test",
		PassWord: "test",
		Email:    "test@test.com",
		Age:      20,
	}

	testPlayerProfile = &models.PlayerProfile{
		Nickname:   "test",
		Avatar:     "test.png",
		Level:      1,
		Experience: 10,
		Points:     100,
		UserID:     testUser.UserID,
	}
)

func TestPlayerProfileRespositoryImpl_CreatePlayerProfile(t *testing.T) {

	t.Run("CreatePlayerProfile_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		// Attempt to create user profile
		require.NoError(t, userRepo.CreateUser(testUser), "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.UserID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, playerRepo.CreatePlayerProfile(testPlayerProfile), "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.PlayerProfileID, "Player Profile ID is zero")
		var dbPlayer models.PlayerProfile
		require.NoError(t, db.Preload("User").First(&dbPlayer, "nickname = ?", testPlayerProfile.Nickname).Error, "Error getting player profile")

		// Assertion
		require.Equal(t, testPlayerProfile.Nickname, dbPlayer.Nickname, "Nicknames do not match")

	})

	t.Run("CreatePlayerProfile_Duplicate", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		// Attempt to create user profile
		require.NoError(t, userRepo.CreateUser(testUser), "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.UserID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, playerRepo.CreatePlayerProfile(testPlayerProfile), "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.PlayerProfileID, "Player Profile ID is zero")

		// Attempt to create Player again
		err := playerRepo.CreatePlayerProfile(testPlayerProfile)
		require.Error(t, err, "Expected error creating duplicate player profile")
	})

}

func TestPlayerProfileRespositoryImpl_GetPlayerProfile(t *testing.T) {

	t.Run("GetPlayerProfile_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		// Attempt to create user profile
		require.NoError(t, userRepo.CreateUser(testUser), "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.UserID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, playerRepo.CreatePlayerProfile(testPlayerProfile), "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.PlayerProfileID, "Player Profile ID is zero")

		// Attempt to get player profile
		dbPlayer, err := playerRepo.GetPlayerProfile(testPlayerProfile.PlayerProfileID)
		require.NoError(t, err, "Error getting player profile")

		// Assertion
		require.Equal(t, testPlayerProfile.Nickname, dbPlayer.Nickname, "Nicknames do not match")
	})

	t.Run("GetPlayerProfile_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)

		// Attempt to get player profile
		_, err := playerRepo.GetPlayerProfile(1)
		require.Error(t, err, "Expected error getting player profile")
	})

}

func TestPlayerProfileRespositoryImpl_UpdatePlayerProfile(t *testing.T) {

	t.Run("UpdatePlayerProfile_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		// Attempt to create user profile
		require.NoError(t, userRepo.CreateUser(testUser), "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.UserID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, playerRepo.CreatePlayerProfile(testPlayerProfile), "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.PlayerProfileID, "Player Profile ID is zero")

		// Update Player
		testPlayerProfile.Nickname = "newNickname"
		require.NoError(t, playerRepo.UpdatePlayerProfile(testPlayerProfile), "Error updating player profile")

		// Get Player
		dbPlayer, err := playerRepo.GetPlayerProfile(testPlayerProfile.PlayerProfileID)
		require.NoError(t, err, "Error getting player profile")

		// Assertion
		require.Equal(t, testPlayerProfile.Nickname, dbPlayer.Nickname, "Nicknames do not match")
	})

	t.Run("UpdatePlayerProfile_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)

		// Attempt to update player profile
		err := playerRepo.UpdatePlayerProfile(testPlayerProfile)
		require.Error(t, err, "Expected error updating player profile")
	})

}

func TestPlayerProfileRespositoryImpl_DeletePlayerProfile(t *testing.T) {

	t.Run("DeletePlayerProfile_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		// Attempt to create user profile
		require.NoError(t, userRepo.CreateUser(testUser), "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.UserID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, playerRepo.CreatePlayerProfile(testPlayerProfile), "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.PlayerProfileID, "Player Profile ID is zero")

		// Attempt to delete player profile
		require.NoError(t, playerRepo.DeletePlayerProfile(testPlayerProfile.PlayerProfileID), "Error deleting player profile")

		// Attempt to get player profile
		_, err := playerRepo.GetPlayerProfile(testPlayerProfile.PlayerProfileID)
		require.Error(t, err, "Expected error getting player profile")
	})

	t.Run("DeletePlayerProfile_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)

		// Attempt to delete player profile
		err := playerRepo.DeletePlayerProfile(1)
		require.Error(t, err, "Expected error deleting player profile")
	})

}

func TestPlayerProfileRespositoryImpl_CheckPlayerProfileExists(t *testing.T) {
	t.Run("CheckPlayerProfileExists_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		// Attempt to create user profile
		require.NoError(t, userRepo.CreateUser(testUser), "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.UserID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, playerRepo.CreatePlayerProfile(testPlayerProfile), "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.PlayerProfileID, "Player Profile ID is zero")

		// Check if Player exists
		exists, err := playerRepo.CheckPlayerProfileExists(testPlayerProfile.PlayerProfileID)
		require.NoError(t, err, "Error checking if player exists")
		require.True(t, exists, "Player does not exist")
	})

	t.Run("CheckPlayerProfileExists_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)

		// Check if Player exists
		exists, err := playerRepo.CheckPlayerProfileExists(1)
		require.NoError(t, err, "Error checking if player exists")
		require.False(t, exists, "Player exists")
	})
}
