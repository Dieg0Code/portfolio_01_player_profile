package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidationPlayerProfile(t *testing.T) {
	t.Run("Validate_Success", func(t *testing.T) {
		validProfile := PlayerProfile{
			Nickname:   "PlayerOne",
			Avatar:     "avatarURL",
			Level:      10,
			Experience: 1000,
			Points:     500,
			UserID:     1,
			User: User{
				UserName: "ValidUser",
				PassWord: "ValidPass123",
				Email:    "user@example.com",
				Age:      30,
				Role:     "user",
			},
		}

		err := validProfile.Validate()
		require.NoError(t, err, "Error validating player profile")
	})

	t.Run("Validate_Invalid", func(t *testing.T) {

		invalidProfile := PlayerProfile{
			Nickname: "PlayerOne",
			Avatar:   "avatarURL",
			Level:    10,
			Points:   500,
		}

		err := invalidProfile.Validate()
		require.Error(t, err, "Expected error validating a invalid player")
	})
}
