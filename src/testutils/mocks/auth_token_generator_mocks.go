package mocks

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

type MockAuthUtils struct {
	mock.Mock
}

func (_m *MockAuthUtils) GenerateToken(userID uint, role string) (string, error) {
	ret := _m.Called(userID, role)
	return ret.String(0), ret.Error(1)
}
func (_m *MockAuthUtils) ParseToken(tokenString string) (*jwt.Token, error) {
	ret := _m.Called(tokenString)
	return ret.Get(0).(*jwt.Token), ret.Error(1)
}
