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

func TestPostUser(t *testing.T) {
	router := gin.Default()
	userRepositoryMock := new(repository.MockUser)
	cldPkgMock := new(pkg.MockCloudinary)

	handler := NewUser(userRepositoryMock, cldPkgMock)

	userRepositoryMock.On("CreateUser", mock.Anything).Return("Data created", nil)
	router.POST("/user", handler.PostUser)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"first_name": "John",
		"last_name": "Doe",
		"phone": "123456789",
		"address": "123 Main St",
		"birth_date": "2000-01-01",
		"email": "test@example.com",
		"password": "testpassword",
		"role":     "user",
	})

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(requestBody))
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

func TestGetUsers(t *testing.T) {
	router := gin.Default()
	userRepositoryMock := new(repository.MockUser)
	cldPkgMock := new(pkg.MockCloudinary)
	hashedPassword, err := pkg.HashPassword("testpassword")
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}
	handler := NewUser(userRepositoryMock, cldPkgMock)

	firstName := "John"
	lastName := "Doe"
	address := "123 Main St"
	userImage := "image.jpg"
	createdAt := time.Now()
	updatedAt := time.Now()

	mockUsers := models.GetUsers{
		{
			User_id:   	   "1",
			User_uuid:	   "uuid1",
			First_name:    &firstName,
			Last_name:     &lastName,
			Phone:         "123456789",
			Address:       &address,
			Email:         "john@example.com",
			Password:      hashedPassword,
			Role:	       "user",
			User_image:    &userImage,
			Created_at:    &createdAt,
			Updated_at:    &updatedAt,
		},
	}

	userRepositoryMock.On("GetAllUser", mock.Anything).Return(&mockUsers, nil)
	router.GET("/user/datauser", handler.GetUsers)

	req, err := http.NewRequest("GET", "/user/datauser?page=1", nil)
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

func TestGetUserDetail(t *testing.T) {
	router := gin.Default()
	userRepositoryMock := new(repository.MockUser)
	cldPkgMock := new(pkg.MockCloudinary)
	hashedPassword, err := pkg.HashPassword("testpassword")
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}
	handler := NewUser(userRepositoryMock, cldPkgMock)

	firstName := "John"
	lastName := "Doe"
	address := "123 Main St"
	userImage := "image.jpg"
	createdAt := time.Now()
	updatedAt := time.Now()

	mockUser := &models.GetUser{
		User_id:   	   "1",
		User_uuid:	   "uuid1",
		First_name:    &firstName,
		Last_name:     &lastName,
		Phone:         "123456789",
		Address:       &address,
		Email:         "john@example.com",
		Password:      hashedPassword,
		Role:	       "user",
		User_image:    &userImage,
		Created_at:    &createdAt,
		Updated_at:    &updatedAt,
	}

	userRepositoryMock.On("GetDetailUser", mock.Anything).Return(mockUser, nil)
	router.GET("/user/profile", func(ctx *gin.Context) {
		ctx.Set("user_id", "1")
		handler.GetUserDetail(ctx)
	})

	req, err := http.NewRequest("GET", "/user/profile", nil)
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

func TestUserDelete(t *testing.T) {
	router := gin.Default()
	userRepositoryMock := new(repository.MockUser)
	cldPkgMock := new(pkg.MockCloudinary)

	handler := NewUser(userRepositoryMock, cldPkgMock)

	userRepositoryMock.On("DeleteUser", mock.Anything).Return("Data deleted", nil)
	router.DELETE("/user/delete", func(ctx *gin.Context) {
		ctx.Set("user_id", "1")
		handler.UserDelete(ctx)
	})

	req, err := http.NewRequest("DELETE", "/user/delete", nil)
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