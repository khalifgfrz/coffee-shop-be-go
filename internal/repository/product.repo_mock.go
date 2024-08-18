package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockProduct struct {
	mock.Mock
}

func (m *MockProduct) CreateProduct(data *models.Product) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockProduct) GetAllProduct(que *models.ProductQuery) (*models.GetProducts, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.GetProducts), args.Error(1)
}

func (m *MockProduct) GetDetailProduct(id string) (*models.GetProduct, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.GetProduct), args.Error(1)
}

func (m *MockProduct) DeleteProduct(id string) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockProduct) UpdateProduct(data *models.Product, id string) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}