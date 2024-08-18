package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockUser struct {
	mock.Mock
}

func (m *MockUser) CreateUser(data *models.User) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUser) GetAllUser(que *models.UserQuery) (*models.GetUsers, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.GetUsers), args.Error(1)
}

func (m *MockUser) GetDetailUser(id string) (*models.GetUser, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.GetUser), args.Error(1)
}

func (m *MockUser) DeleteUser(id string) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUser) UpdateUser(data *models.User, id string) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}