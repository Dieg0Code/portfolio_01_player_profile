package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidationUser(t *testing.T) {
	t.Run("Validate_Success", func(t *testing.T) {
		validUser := User{

			UserName: "validUser",
			PassWord: "validPass",
			Email:    "valid@example.com",
			Age:      25,
		}

		err := validUser.Validate()
		require.NoError(t, err, "Error validating user")
	})

	t.Run("Validate_Invalid", func(t *testing.T) {

		invalidUser := User{
			PassWord: "validPass",
			Email:    "valid@example.com",
		}

		err := invalidUser.Validate()
		require.Error(t, err, "Error validating user")
	})
}
