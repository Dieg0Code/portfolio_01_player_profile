package impl

import (
	"github.com/dieg0code/player-profile/src/services"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func NewPassWordHasher() services.PasswordHasher {
	return &BcryptHasher{}
}

func (h *BcryptHasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Error("[BcryptHasher.HashPassword] Failed to hash password")
		return "", err
	}

	return string(hashedPassword), nil
}

func (h *BcryptHasher) ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
