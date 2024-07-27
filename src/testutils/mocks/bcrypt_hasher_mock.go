package mocks

import "github.com/stretchr/testify/mock"

type MockPasswordHasher struct {
	mock.Mock
}

func (_m *MockPasswordHasher) HashPassword(password string) (string, error) {
	args := _m.Called(password)
	return args.String(0), args.Error(1)
}

func (_m *MockPasswordHasher) ComparePassword(hashedPassword string, password string) error {
	args := _m.Called(hashedPassword, password)
	return args.Error(0)
}
