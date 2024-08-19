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

func TestPostFavorite(t *testing.T) {
	router := gin.Default()
	favoriteRepositoryMock := new(repository.MockFavorite)

	handler := NewFavorite(favoriteRepositoryMock)

	favoriteRepositoryMock.On("CreateFavorite", mock.Anything).Return("Data created", nil)
	router.POST("/favorite", handler.PostFavorite)

	requestBody, _ := json.Marshal(map[string]int{
		"product_id": 1,
		"user_id":    1,
	})

	req, err := http.NewRequest("POST", "/favorite", bytes.NewBuffer(requestBody))
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

func TestGetFavorites(t *testing.T) {
	router := gin.Default()
	favoriteRepositoryMock := new(repository.MockFavorite)
	handler := NewFavorite(favoriteRepositoryMock)

	firstName := "John"
	lastName := "Doe"
	phone := "123456789"
	address := "123 Main St"
	email := "john@example.com"
	userImage := "image.jpg"
	productName := "Coffee"
	price := "5000"
	category := "Beverage"
	description := "A hot coffee"
	productImage := "coffee.jpg"
	createdAt := time.Now()
	updatedAt := time.Now()

	mockFavorites := models.GetFavorites{
		{
			Favorite_id:   "1",
			Favorite_uuid: "uuid1",
			First_name:    &firstName,
			Last_name:     &lastName,
			Phone:         &phone,
			Address:       &address,
			Email:         &email,
			User_image:    &userImage,
			Product_name:  &productName,
			Price:         &price,
			Category:      &category,
			Description:   &description,
			Product_image: &productImage,
			Created_at:    &createdAt,
			Updated_at:    &updatedAt,
		},
	}

	favoriteRepositoryMock.On("GetAllFavorite", mock.Anything).Return(&mockFavorites, nil)
	router.GET("/favorite", handler.GetFavorites)

	req, err := http.NewRequest("GET", "/favorite?page=1", nil)
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

func TestGetFavoriteDetail(t *testing.T) {
	router := gin.Default()
	favoriteRepositoryMock := new(repository.MockFavorite)
	handler := NewFavorite(favoriteRepositoryMock)

	firstName := "John"
	lastName := "Doe"
	phone := "123456789"
	address := "123 Main St"
	email := "john@example.com"
	userImage := "image.jpg"
	productName := "Coffee"
	price := "5000"
	category := "Beverage"
	description := "A hot coffee"
	productImage := "coffee.jpg"
	createdAt := time.Now()
	updatedAt := time.Now()

	mockFavorite := &models.GetFavorite{
		Favorite_id:   "1",
		Favorite_uuid: "uuid1",
		First_name:    &firstName,
		Last_name:     &lastName,
		Phone:         &phone,
		Address:       &address,
		Email:         &email,
		User_image:    &userImage,
		Product_name:  &productName,
		Price:         &price,
		Category:      &category,
		Description:   &description,
		Product_image: &productImage,
		Created_at:    &createdAt,
		Updated_at:    &updatedAt,
	}

	favoriteRepositoryMock.On("GetDetailFavorite", mock.Anything).Return(mockFavorite, nil)
	router.GET("/favorite/:id", handler.GetFavoriteDetail)

	req, err := http.NewRequest("GET", "/favorite/1", nil)
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

func TestFavoriteDelete(t *testing.T) {
	router := gin.Default()
	favoriteRepositoryMock := new(repository.MockFavorite)

	handler := NewFavorite(favoriteRepositoryMock)
	favoriteRepositoryMock.On("DeleteFavorite", mock.Anything).Return("Data deleted", nil)
	router.DELETE("/favorite/:id", handler.FavoriteDelete)

	req, err := http.NewRequest("DELETE", "/favorite/1", nil)
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