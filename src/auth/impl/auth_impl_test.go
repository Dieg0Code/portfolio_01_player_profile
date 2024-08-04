package impl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthImpl(t *testing.T) {
	t.Run("GenrateToken", func(t *testing.T) {
		auth := AuthImpl{}
		token, err := auth.GenerateToken(
			1,
			"admin",
		)

		assert.Nil(t, err, "Expected no error generating token")
		assert.NotEmpty(t, token, "Expected token to be generated")
	})
}
