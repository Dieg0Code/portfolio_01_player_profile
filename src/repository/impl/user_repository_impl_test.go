package impl

import (
	"testing"

	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/testutils"
	"github.com/stretchr/testify/require"
)

func TestUserRepositoryImpl_CreateUser(t *testing.T) {

	// Setup
	db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{})
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create repository instance
	repo := NewUserRepositoryImpl(db)

	// Define user to be created
	user := &models.User{
		UserName: "test",
		PassWord: "test",
		Email:    "test@test.com",
		Age:      20,
		Profiles: []models.PlayerProfile{},
	}

	// Attempt to create user
	require.NoError(t, repo.CreateUser(user), "Error creating user")

	// Verify user was created
	var dbUser models.User
	require.NoError(t, db.Preload("Profiles").First(&dbUser, "user_name = ?", user.UserName).Error, "Error getting user")

	// Assertions
	require.Equal(t, user.UserName, dbUser.UserName, "Usernames do not match")
}

func TestUserRepositoryImpl_GetUser(t *testing.T) {

	// Setup
	db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{})
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create repository instance
	repo := NewUserRepositoryImpl(db)

	// Define user to be created
	user := &models.User{
		UserName: "test",
		PassWord: "test",
		Email:    "test@test.com",
		Age:      20,
		Profiles: []models.PlayerProfile{},
	}

	// Attempt to create user
	require.NoError(t, repo.CreateUser(user), "Error creating user")

	// Verify user was created
	require.NotZero(t, user.UserID, "User ID is zero")

	// Attempt to get user
	dbUser, err := repo.GetUser(user.UserID)
	require.NoError(t, err, "Error getting user")

	// Assertions
	require.Equal(t, user.UserName, dbUser.UserName, "Usernames do not match")
}

func TestUserRepositoryImpl_UpdateUser(t *testing.T) {

	// Setup
	db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{})
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create repository instance
	repo := NewUserRepositoryImpl(db)

	// Define user to be created
	user := &models.User{
		UserName: "test",
		PassWord: "test",
		Email:    "test@test.com",
		Age:      20,
		Profiles: []models.PlayerProfile{},
	}

	// Attempt to create user
	require.NoError(t, repo.CreateUser(user), "Error creating user")

	// Update user
	user.Email = "UPDATED"

	// Attempt to update user
	require.NoError(t, repo.UpdateUser(user), "Error updating user")

	// Verify user was updated
	var dbUser models.User
	require.NoError(t, db.Preload("Profiles").First(&dbUser, "user_name = ?", user.UserName).Error, "Error getting user")

	// Assertions
	require.Equal(t, user.Email, dbUser.Email, "Emails do not match")
}

func TestUserRepositoryImpl_DeleteUser(t *testing.T) {

	// Setup
	db := testutils.SetupTestDB(&models.User{}, &models.PlayerProfile{})
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create repository instance
	repo := NewUserRepositoryImpl(db)

	// Define user to be created
	user := &models.User{
		UserName: "test",
		PassWord: "test",
		Email:    "test@test.com",
		Age:      20,
		Profiles: []models.PlayerProfile{},
	}

	// Attempt to create user
	require.NoError(t, repo.CreateUser(user), "Error creating user")

	// Verify user was created
	require.NotZero(t, user.UserID, "User ID is zero")

	// Attempt to delete user
	require.NoError(t, repo.DeleteUser(user.UserID), "Error deleting user")

	// Verify user was deleted
	var dbUser models.User
	require.Error(t, db.Preload("Profiles").First(&dbUser, "user_name = ?", user.UserName).Error, "Error getting user")

	// Assertions
	require.Equal(t, models.User{}, dbUser, "User is not empty")
}
