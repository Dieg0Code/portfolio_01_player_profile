package impl

import (
	"testing"

	"github.com/dieg0code/player-profile/src/helpers"
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
		UserID:     testUser.ID,
	}
)

func TestNewPlayerProfileRepositoryImpl(t *testing.T) {
	db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
	defer func() {
		sqlDB, _ := db.DB()
		err := sqlDB.Close()
		if err != nil {
			t.Errorf("Error closing database connection: %v", err)
		}
	}()

	playerRepo := NewPlayerProfileRepositoryImpl(db)
	require.NotNil(t, playerRepo, "PlayerProfileRepositoryImpl is nil")
}

func TestPlayerProfileRespositoryImpl_CreatePlayerProfile(t *testing.T) {

	t.Run("CreatePlayerProfile_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		resultUser := userRepo.CreateUser(testUser)
		resultPlayer := playerRepo.CreatePlayerProfile(testPlayerProfile)

		// Attempt to create user profile
		require.NoError(t, resultUser, "Error creating user")
		require.Nil(t, resultUser, "Error creating user")
		// Verify user was created
		require.NotZero(t, testUser.ID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, resultPlayer, "Error creating player profile")
		require.Nil(t, resultPlayer, "Error creating player profile")
		// Varify Player was created
		require.NotZero(t, testPlayerProfile.ID, "Player Profile ID is zero")

		var dbPlayer models.PlayerProfile
		require.NoError(t, db.Preload("User").First(&dbPlayer, "nickname = ?", testPlayerProfile.Nickname).Error, "Error getting player profile")

		// Assertion
		require.Equal(t, testPlayerProfile.Nickname, dbPlayer.Nickname, "Nicknames do not match")

	})

	t.Run("CreatePlayerProfile_Duplicate", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		resultUser := userRepo.CreateUser(testUser)
		resultPlayer := playerRepo.CreatePlayerProfile(testPlayerProfile)

		// Attempt to create user profile
		require.NoError(t, resultUser, "Error creating user")
		require.Nil(t, resultUser, "Error creating user")
		// Verify user was created
		require.NotZero(t, testUser.ID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, resultPlayer, "Error creating player profile")
		require.Nil(t, resultPlayer, "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.ID, "Player Profile ID is zero")

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
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		resultUser := userRepo.CreateUser(testUser)
		resultPlayer := playerRepo.CreatePlayerProfile(testPlayerProfile)

		// Attempt to create user profile
		require.NoError(t, resultUser, "Error creating user")
		require.Nil(t, resultUser, "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.ID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, resultPlayer, "Error creating player profile")
		require.Nil(t, resultPlayer, "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.ID, "Player Profile ID is zero")

		// Attempt to get player profile
		dbPlayer, err := playerRepo.GetPlayerProfile(testPlayerProfile.ID)
		require.NoError(t, err, "Error getting player profile")

		// Assertion
		require.Equal(t, testPlayerProfile.Nickname, dbPlayer.Nickname, "Nicknames do not match")
	})

	t.Run("GetPlayerProfile_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)

		// Attempt to get player profile
		_, err := playerRepo.GetPlayerProfile(1)
		require.Error(t, err, "Expected error getting player profile")
		require.EqualError(t, err, helpers.ErrorPlayerProfileNotFound.Error(), "Error messages do not match")
	})

}

func TestPlayerProfileRespositoryImpl_UpdatePlayerProfile(t *testing.T) {

	t.Run("UpdatePlayerProfile_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		resultUser := userRepo.CreateUser(testUser)
		resultPlayer := playerRepo.CreatePlayerProfile(testPlayerProfile)

		// Attempt to create user profile
		require.NoError(t, resultUser, "Error creating user")
		require.Nil(t, resultUser, "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.ID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, resultPlayer, "Error creating player profile")
		require.Nil(t, resultPlayer, "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.ID, "Player Profile ID is zero")

		// Update Player
		testPlayerProfile.Nickname = "newNickname"
		resultUpdatePlayer := playerRepo.UpdatePlayerProfile(testPlayerProfile.ID, testPlayerProfile)
		require.NoError(t, resultUpdatePlayer, "Error updating player profile")

		// Get Player
		dbPlayer, err := playerRepo.GetPlayerProfile(testPlayerProfile.ID)
		require.NoError(t, err, "Error getting player profile")

		// Assertion
		require.Equal(t, testPlayerProfile.Nickname, dbPlayer.Nickname, "Nicknames do not match")
	})

	t.Run("UpdatePlayerProfile_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)

		// Attempt to update player profile
		err := playerRepo.UpdatePlayerProfile(testPlayerProfile.ID, testPlayerProfile)
		require.Error(t, err, "Expected error updating player profile")
		require.EqualError(t, err, helpers.ErrorPlayerProfileNotFound.Error(), "Error messages do not match")
	})

}

func TestPlayerProfileRespositoryImpl_DeletePlayerProfile(t *testing.T) {

	t.Run("DeletePlayerProfile_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		resultCreateUser := userRepo.CreateUser(testUser)
		resultCreatePlayer := playerRepo.CreatePlayerProfile(testPlayerProfile)

		// Attempt to create user profile
		require.NoError(t, resultCreateUser, "Error creating user")
		require.Nil(t, resultCreateUser, "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.ID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, resultCreatePlayer, "Error creating player profile")
		require.Nil(t, resultCreatePlayer, "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.ID, "Player Profile ID is zero")

		resultDeletePlayer := playerRepo.DeletePlayerProfile(testPlayerProfile.ID)

		// Attempt to delete player profile
		require.NoError(t, resultDeletePlayer, "Error deleting player profile")

		// Attempt to get player profile
		_, err := playerRepo.GetPlayerProfile(testPlayerProfile.ID)
		require.Error(t, err, "Expected error getting player profile")
	})

	t.Run("DeletePlayerProfile_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)

		// Attempt to delete player profile
		err := playerRepo.DeletePlayerProfile(1)
		require.Error(t, err, "Expected error deleting player profile")
		require.EqualError(t, err, helpers.ErrorPlayerProfileNotFound.Error(), "Error messages do not match")
	})

}

func TestPlayerProfileRespositoryImpl_CheckPlayerProfileExists(t *testing.T) {
	t.Run("CheckPlayerProfileExists_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)
		userRepo := NewUserRepositoryImpl(db)

		resultCreateUser := userRepo.CreateUser(testUser)
		resultCreatePlayer := playerRepo.CreatePlayerProfile(testPlayerProfile)

		// Attempt to create user profile
		require.NoError(t, resultCreateUser, "Error creating user")
		require.Nil(t, resultCreateUser, "Error creating user")

		// Verify user was created
		require.NotZero(t, testUser.ID, "User ID is zero")

		// Attempt to create Player
		require.NoError(t, resultCreatePlayer, "Error creating player profile")
		require.Nil(t, resultCreatePlayer, "Error creating player profile")

		// Varify Player was created
		require.NotZero(t, testPlayerProfile.ID, "Player Profile ID is zero")

		// Check if Player exists
		exists, err := playerRepo.CheckPlayerProfileExists(testPlayerProfile.ID)
		require.NoError(t, err, "Error checking if player exists")
		require.True(t, exists, "Player does not exist")
	})

	t.Run("CheckPlayerProfileExists_NotFound", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.PlayerProfile{}, &models.User{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Errorf("Error closing database connection: %v", err)
			}
		}()

		playerRepo := NewPlayerProfileRepositoryImpl(db)

		// Check if Player exists
		exists, err := playerRepo.CheckPlayerProfileExists(1)
		require.NoError(t, err, "Error checking if player exists")
		require.False(t, exists, "Player exists")
	})
}
