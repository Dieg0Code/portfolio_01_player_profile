package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidationAchievement(t *testing.T) {
	t.Run("Validate_Success", func(t *testing.T) {
		validAchievement := Achievement{
			Name:        "First Achievement",
			Description: "Description of the first achievement",
		}

		err := validAchievement.Validate()
		require.NoError(t, err, "Error validating achievement")
	})

	t.Run("Validate_Invalid", func(t *testing.T) {

		invalidAchievement := Achievement{
			Name: "First Achievement",
		}

		err := invalidAchievement.Validate()
		require.Error(t, err, "Error validating achievement")
	})
}
