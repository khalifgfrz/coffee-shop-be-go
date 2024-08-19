package handlers

import (
	"bytes"
	"encoding/json"
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostProduct(t *testing.T) {
	router := gin.Default()
	productRepositoryMock := new(repository.MockProduct)
	cldPkgMock := new(pkg.MockCloudinary)

	handler := NewProduct(productRepositoryMock, cldPkgMock)

	productRepositoryMock.On("CreateProduct", mock.Anything).Return("Data created", nil)
	router.POST("/product", handler.PostProduct)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"product_name": "Coffee",
		"price": 0,
		"category": "Brewed",
		"description": "example",
		"stock": 0,
	})

	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, 201, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Create data success", actualResponse.Message, "Response message does not match")
	assert.Equal(t, "Data created", actualResponse.Data, "Data does not match")
}

func TestGetProducts(t *testing.T) {
	router := gin.Default()
	productRepositoryMock := new(repository.MockProduct)
	cldPkgMock := new(pkg.MockCloudinary)

	handler := NewProduct(productRepositoryMock, cldPkgMock)

	productName := "Coffee"
	price := 0
	category := "Brewed"
	description := "example"
	stock := 0
	productImage := "image.jpg"
	createdAt := time.Now()
	updatedAt := time.Now()

	mockProducts := models.GetProducts{
		{
			Product_id:   "1",
			Product_uuid: "uuid1",
			Product_name:  &productName,
			Price:         &price,
			Category:      &category,
			Description:   &description,
			Stock:   	   &stock,
			Product_image: &productImage,
			Created_at:    &createdAt,
			Updated_at:    &updatedAt,
		},
	}

	productRepositoryMock.On("GetAllProduct", mock.Anything).Return(&mockProducts, nil)
	router.GET("/product", handler.GetProducts)

	req, err := http.NewRequest("GET", "/product?page=1", nil)
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")
	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, http.StatusOK, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Get data success", actualResponse.Message, "Response message does not match")
	assert.NotNil(t, actualResponse.Data, "Data does not match")
}

func TestGetProductDetail(t *testing.T) {
	router := gin.Default()
	productRepositoryMock := new(repository.MockProduct)
	cldPkgMock := new(pkg.MockCloudinary)

	handler := NewProduct(productRepositoryMock, cldPkgMock)

	productName := "Coffee"
	price := 0
	category := "Brewed"
	description := "example"
	stock := 0
	productImage := "image.jpg"
	createdAt := time.Now()
	updatedAt := time.Now()

	mockProduct := &models.GetProduct{
		Product_id:   "1",
		Product_uuid: "uuid1",
		Product_name:  &productName,
		Price:         &price,
		Category:      &category,
		Description:   &description,
		Stock:   	   &stock,
		Product_image: &productImage,
		Created_at:    &createdAt,
		Updated_at:    &updatedAt,
	}

	productRepositoryMock.On("GetDetailProduct", mock.Anything).Return(mockProduct, nil)
	router.GET("/product/:id", handler.GetProductDetail)

	req, err := http.NewRequest("GET", "/product/1", nil)
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")
	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, http.StatusOK, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Get data success", actualResponse.Message, "Response message does not match")
	assert.NotNil(t, actualResponse.Data, "Data does not match")
}

func TestProductDelete(t *testing.T) {
	router := gin.Default()
	productRepositoryMock := new(repository.MockProduct)
	cldPkgMock := new(pkg.MockCloudinary)

	handler := NewProduct(productRepositoryMock, cldPkgMock)

	productRepositoryMock.On("DeleteProduct", mock.Anything).Return("Data deleted", nil)
	router.DELETE("/product/:id", handler.ProductDelete)

	req, err := http.NewRequest("DELETE", "/product/1", nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, 200, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Delete data success", actualResponse.Message, "Response message does not match")
	assert.Equal(t, "Data deleted", actualResponse.Data, "Status does not match")
}