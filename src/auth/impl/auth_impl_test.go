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

func TestAuthImpl_ParseToken(t *testing.T) {
	t.Run("ValidToken", func(t *testing.T) {
		auth := AuthImpl{}
		token, err := auth.GenerateToken(
			1,
			"admin",
		)
		assert.Nil(t, err, "Expected no error generating token")

		_, err = auth.ParseToken(token)
		assert.Nil(t, err, "Expected no error parsing token")
	})

	t.Run("InvalidToken", func(t *testing.T) {
		auth := AuthImpl{}
		_, err := auth.ParseToken("invalidtoken")
		assert.NotNil(t, err, "Expected error parsing token")
	})

}
