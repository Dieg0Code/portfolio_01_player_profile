package mocks

import (
	"github.com/dieg0code/player-profile/src/models"
	"github.com/stretchr/testify/mock"
)

type MockPlayerProfileController struct {
	mock.Mock
}

func (m *MockPlayerProfileController) GetPlayerByIDFromService(id uint) (*models.PlayerProfile, error) {
	args := m.Called(id)
	return args.Get(0).(*models.PlayerProfile), args.Error(1)
}
