package impl

import (
	"github.com/dieg0code/player-profile/src/models"
	repo "github.com/dieg0code/player-profile/src/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repo.UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// CreateUser implements repository.UserRepository.
func (u *UserRepositoryImpl) CreateUser(user *models.User) error {
	panic("unimplemented")
}

// GetUser implements repository.UserRepository.
func (u *UserRepositoryImpl) GetUser(userID int) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements repository.UserRepository.
func (u *UserRepositoryImpl) UpdateUser(user *models.User) error {
	panic("unimplemented")
}

// DeleteUser implements repository.UserRepository.
func (u *UserRepositoryImpl) DeleteUser(userID int) error {
	panic("unimplemented")
}
