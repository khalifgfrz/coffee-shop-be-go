package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockFavorite struct {
	mock.Mock
}

func (m *MockFavorite) CreateFavorite(data *models.Favorite) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockFavorite) GetAllFavorite(que *models.FavoriteQuery) (*models.GetFavorites, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.GetFavorites), args.Error(1)
}

func (m *MockFavorite) GetDetailFavorite(id string) (*models.GetFavorite, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.GetFavorite), args.Error(1)
}

func (m *MockFavorite) DeleteFavorite(id string) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}