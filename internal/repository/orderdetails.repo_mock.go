package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockOrderDetails struct {
	mock.Mock
}

func (m *MockOrderDetails) 	CreateOrderDetails(order_id int, products []models.ProductDetail) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockOrderDetails)	GetOrderDetails(order_id int) (*[]models.GetOrderDetail, error) {
	args := m.Mock.Called()
	return args.Get(0).(*[]models.GetOrderDetail), args.Error(1)
}