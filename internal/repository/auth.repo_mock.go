package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockAuth struct {
	mock.Mock
}

func (m *MockAuth) RegisterUser(data *models.Auth) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockAuth) LoginUser(email string) (*models.Auth, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.Auth), args.Error(1)
}