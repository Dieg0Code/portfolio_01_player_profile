package impl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcryptHasher(t *testing.T) {
	hasher := NewPassWordHasher()
	password := "password"

	t.Run("HashPassword_Success", func(t *testing.T) {
		hashedPassword, err := hasher.HashPassword(password)
		assert.Nil(t, err, "Expected no error hashing password")
		assert.NotEmpty(t, hashedPassword, "Expected hashed password not empty")
	})

	t.Run("ComparePassword_Success", func(t *testing.T) {
		hashedPassword, err := hasher.HashPassword(password)
		assert.Nil(t, err, "Expected no error hashing password")
		assert.NotEmpty(t, hashedPassword, "Expected hashed password not empty")

		err = hasher.ComparePassword(hashedPassword, password)
		assert.Nil(t, err, "Expected no error comparing password")
	})

	t.Run("ComparePassword_Fail", func(t *testing.T) {
		hashedPassword, err := hasher.HashPassword(password)
		assert.Nil(t, err, "Expected no error hashing password")
		assert.NotEmpty(t, hashedPassword, "Expected hashed password not empty")

		err = hasher.ComparePassword(hashedPassword, "invalid")
		assert.NotNil(t, err, "Expected error comparing password")
	})

}
