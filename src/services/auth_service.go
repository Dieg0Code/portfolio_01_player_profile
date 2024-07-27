package services

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
)

type AuthService interface {
	Login(loginRequest request.LoginRequest) (*response.LoginResponse, error)
}
