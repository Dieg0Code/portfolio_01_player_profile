package repository

import "github.com/dieg0code/player-profile/src/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(userID uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	UpdateUser(userID uint, user *models.User) error
	DeleteUser(userID uint) error
	GetAllUsers(offset int, pageSize int) ([]models.User, error)
}
