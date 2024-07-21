package services

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest) error
	GetByID(userID uint) (*response.UserResponse, error)
	Update(userID uint, user request.UpdateUserRequest) error
	Delete(userID uint) error
	GetAll() ([]response.UserResponse, error)
}
