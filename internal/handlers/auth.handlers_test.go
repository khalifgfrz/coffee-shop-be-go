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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	router := gin.Default()
	authRepositoryMock := new(repository.MockAuth)

	handler := NewAuth(authRepositoryMock)
	authRepositoryMock.On("RegisterUser", mock.Anything).Return("Data created", nil)
	router.POST("/auth/register", handler.Register)

	requestBody, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "testpassword",
		"role":     "user",
	})

	req, err := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code, "Status code does not match")
	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, 201, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Register success", actualResponse.Message, "Response message does not match")
	assert.Equal(t, "Data created", actualResponse.Data, "Status does not match")
}

func TestLogin(t *testing.T) {
	router := gin.Default()
	authRepositoryMock := new(repository.MockAuth)
	hashedPassword, err := pkg.HashPassword("testpassword")
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	handler := NewAuth(authRepositoryMock)
	mockUser := &models.Auth{
		User_id:  "1",
		Email:    "test@example.com",
		Password: hashedPassword,
		Role:     "user",
	}
	authRepositoryMock.On("LoginUser", mock.Anything).Return(mockUser, nil)
	router.POST("/auth/login", handler.Login)

	requestBody, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "testpassword",
	})

	req, err := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code, "Status code does not match")
	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, http.StatusCreated, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Login success", actualResponse.Message, "Response message does not match")
	assert.NotNil(t, actualResponse.Data, "Token should be returned")
}

