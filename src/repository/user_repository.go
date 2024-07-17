package repository

import "github.com/dieg0code/player-profile/src/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(userID int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID int) error
}
