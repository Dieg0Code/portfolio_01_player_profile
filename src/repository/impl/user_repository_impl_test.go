package impl

import (
	"testing"

	"github.com/dieg0code/player-profile/src/models"
	testutils "github.com/dieg0code/player-profile/src/test"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	testutils.RunTestsWithDatabase(m, &models.User{}, &models.PlayerProfile{})
}

func TestUserRepositoryImpl_CreateUser(t *testing.T) {

	db := testutils.Db

	// Create repository instance
	repo := NewUserRepositoryImpl(db)

	// Define user to be created
	user := &models.User{
		UserName: "test",
		PassWord: "test",
		Email:    "test@test.com",
		Age:      20,
		Profiles: []models.PlayerProfile{
			{
				Nickname:   "testProfile",
				Avatar:     "default.png",
				Level:      1,
				Experience: 0,
				Points:     100,
			},
		},
	}

	// Attempt to create user
	require.NoError(t, repo.CreateUser(user))

	// Verify user was created
	var dbUser models.User
	require.NoError(t, db.Preload("Profiles").First(&dbUser, "user_name = ?", user.UserName).Error)

	// Assertions
	require.Equal(t, user.UserName, dbUser.UserName)
	require.Len(t, dbUser.Profiles, 1)
}

func TestUserRepositoryImpl_GetUser(t *testing.T) {

	db := testutils.Db

	// Create repository instance
	repo := NewUserRepositoryImpl(db)

	// Define user to be created
	user := &models.User{
		UserName: "test",
		PassWord: "test",
		Email:    "test@test.com",
		Age:      20,
		Profiles: []models.PlayerProfile{
			{
				Nickname:   "testProfile",
				Avatar:     "default.png",
				Level:      1,
				Experience: 0,
				Points:     100,
			},
		},
	}

	// Attempt to create user
	require.NoError(t, repo.CreateUser(user))

	// Verify user was created
	require.NotZero(t, user.UserID)

	// Attempt to get user
	dbUser, err := repo.GetUser(user.UserID)
	require.NoError(t, err)

	// Assertions
	require.Equal(t, user.UserName, dbUser.UserName)
	require.Len(t, dbUser.Profiles, 1)
}

func TestUserRepositoryImpl_UpdateUser(t *testing.T) {

	db := testutils.Db

	// Create repository instance
	repo := NewUserRepositoryImpl(db)

	// Define user to be created
	user := &models.User{
		UserName: "test",
		PassWord: "test",
		Email:    "test@test.com",
		Age:      20,
		Profiles: []models.PlayerProfile{
			{
				Nickname:   "testProfile",
				Avatar:     "default.png",
				Level:      1,
				Experience: 0,
				Points:     100,
			},
		},
	}

	// Attempt to create user
	require.NoError(t, repo.CreateUser(user))

	// Update user
	user.Email = "UPDATED"

	// Attempt to update user
	require.NoError(t, repo.UpdateUser(user))

	// Verify user was updated
	var dbUser models.User
	require.NoError(t, db.Preload("Profiles").First(&dbUser, "user_name = ?", user.UserName).Error)

	// Assertions
	require.Equal(t, user.Email, dbUser.Email)
	require.Len(t, dbUser.Profiles, 1)
}

func TestUserRepositoryImpl_DeleteUser(t *testing.T) {

	db := testutils.Db

	// Create repository instance
	repo := NewUserRepositoryImpl(db)

	// Define user to be created
	user := &models.User{
		UserName: "test",
		PassWord: "test",
		Email:    "test@test.com",
		Age:      20,
		Profiles: []models.PlayerProfile{
			{
				Nickname:   "testProfile",
				Avatar:     "default.png",
				Level:      1,
				Experience: 0,
				Points:     100,
			},
		},
	}

	// Attempt to create user
	require.NoError(t, repo.CreateUser(user))

	// Verify user was created
	require.NotZero(t, user.UserID)

	// Attempt to delete user
	require.NoError(t, repo.DeleteUser(user.UserID))

	// Verify user was deleted
	var dbUser models.User
	require.Error(t, db.Preload("Profiles").First(&dbUser, "user_name = ?", user.UserName).Error)

	// Assertions
	require.Equal(t, models.User{}, dbUser)
}
