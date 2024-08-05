package mocks

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (_m *MockAuthService) Login(loginRequest request.LoginRequest) (*response.LoginResponse, error) {
	ret := _m.Called(loginRequest)
	return ret.Get(0).(*response.LoginResponse), ret.Error(1)
}
