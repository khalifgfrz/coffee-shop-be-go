package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockOrder struct {
	mock.Mock
}

func (m *MockOrder) CreateOrder(order *models.Order) (int, error) {
	args := m.Mock.Called()
	return args.Get(0).(int), args.Error(1)
}

func (m *MockOrder) GetAllOrder(que *models.OrderQuery) (*models.GetOrders, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.GetOrders), args.Error(1)
}

func (m *MockOrder) GetDetailOrder(uuid string) (*models.GetOrder, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.GetOrder), args.Error(1)
}

func (m *MockOrder) GetHistoryOrder(id string) (*models.GetOrder, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.GetOrder), args.Error(1)
}

func (m *MockOrder) UpdateOrder(data *models.Order, uuid string) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}