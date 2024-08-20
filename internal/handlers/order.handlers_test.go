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

func TestPostOrder(t *testing.T) {
	router := gin.Default()
	orderRepositoryMock := new(repository.MockOrder)
	orderDetailRepositoryMock := new(repository.MockOrderDetails)
	handler := NewOrder(orderRepositoryMock, orderDetailRepositoryMock)

	order := &models.Order{
		Orderlist_id: 1,
		Orderlist_uuid: "uuid1",
		User_id:      1,
		Subtotal:      1,
		Tax:      1,
		Payment_id:      1,
		Delivery_id:      1,
		Status:       "pending",
		Grand_total:   100000,
		Products: []models.ProductDetail{
			{
				Size_id: 1,
				Product_id: 1,
				Qty:    2,
			},
		},
	}
	
	orderRepositoryMock.On("CreateOrder", mock.Anything).Return(order.Orderlist_id, nil)
	orderDetailRepositoryMock.On("CreateOrderDetails", mock.Anything).Return("Data created", nil)
	router.POST("/order", handler.PostOrder)

	requestBody, _ := json.Marshal(order)

	req, err := http.NewRequest("POST", "/order", bytes.NewBuffer(requestBody))
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

func TestGetOrders(t *testing.T) {
	router := gin.Default()
	orderRepositoryMock := new(repository.MockOrder)
	orderDetailRepositoryMock := new(repository.MockOrderDetails)
	handler := NewOrder(orderRepositoryMock, orderDetailRepositoryMock)

	firstName := "John"
	lastName := "Doe"
	phone := "123456789"
	address := "123 Main St"
	email := "john@example.com"
	subtotal := "0"
	tax := "0"
	paymentMethod := "Cash"
	deliveryMethod := "Dine in"
	addedCost := "0"
	status := "waiting"
	grandTotal := "0"
	orderID := 1
	productImage := "image.jpg"
	productName := "Coffee"
	price := 0
	size := "R"
	cost := 0
	qty := 1

	mockOrders := models.GetOrders{
		{
			Orderlist_id:   1,
			Orderlist_uuid: "uuid1",
			First_name:      &firstName,
			Last_name:       &lastName,
			Phone:           &phone,
			Address:         &address,
			Email:           &email,
			Subtotal: 	     &subtotal,
			Tax: 		     &tax,
			Payment_method:  &paymentMethod,
			Delivery_method: &deliveryMethod,
			Added_cost:	     &addedCost,
			Status:			 &status,
			Grand_total:     &grandTotal,
			Products: []models.GetOrderDetail{
				{
					Order_id: &orderID,
					Product_image: &productImage,
					Product_name:    &productName,
					Price: &price,
					Size: &size,
					Added_cost: &cost,
					Qty: &qty,
				},
			},
		},
	}

	orderRepositoryMock.On("GetAllOrder", mock.Anything).Return(&mockOrders, nil)
	orderDetailRepositoryMock.On("GetOrderDetails", mock.Anything).Return(&mockOrders[0].Products, nil)
	router.GET("/order", handler.GetOrders)

	req, err := http.NewRequest("GET", "/order?page=1", nil)
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

func TestGetOrderDetail(t *testing.T) {
	router := gin.Default()
	orderRepositoryMock := new(repository.MockOrder)
	orderDetailRepositoryMock := new(repository.MockOrderDetails)
	handler := NewOrder(orderRepositoryMock, orderDetailRepositoryMock)

	firstName := "John"
	lastName := "Doe"
	phone := "123456789"
	address := "123 Main St"
	email := "john@example.com"
	subtotal := "0"
	tax := "0"
	paymentMethod := "Cash"
	deliveryMethod := "Dine in"
	addedCost := "0"
	status := "waiting"
	grandTotal := "0"
	orderID := 1
	productImage := "image.jpg"
	productName := "Coffee"
	price := 0
	size := "R"
	cost := 0
	qty := 1

	mockOrder := models.GetOrder{
			Orderlist_id:   1,
			Orderlist_uuid: "uuid1",
			First_name:      &firstName,
			Last_name:       &lastName,
			Phone:           &phone,
			Address:         &address,
			Email:           &email,
			Subtotal: 	     &subtotal,
			Tax: 		     &tax,
			Payment_method:  &paymentMethod,
			Delivery_method: &deliveryMethod,
			Added_cost:	     &addedCost,
			Status:			 &status,
			Grand_total:     &grandTotal,
			Products: []models.GetOrderDetail{
				{
					Order_id: &orderID,
					Product_image: &productImage,
					Product_name:    &productName,
					Price: &price,
					Size: &size,
					Added_cost: &cost,
					Qty: &qty,
				},
			},
		}

	orderRepositoryMock.On("GetDetailOrder", mock.Anything).Return(&mockOrder, nil)
	orderDetailRepositoryMock.On("GetOrderDetails", mock.Anything).Return(&mockOrder.Products, nil)
	router.GET("/order/:uuid", handler.GetOrderDetail)
	

	req, err := http.NewRequest("GET", "/order/uuid1", nil)
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

func TestGetOrderHistory(t *testing.T) {
	router := gin.Default()
	orderRepositoryMock := new(repository.MockOrder)
	orderDetailRepositoryMock := new(repository.MockOrderDetails)
	handler := NewOrder(orderRepositoryMock, orderDetailRepositoryMock)

	firstName := "John"
	lastName := "Doe"
	phone := "123456789"
	address := "123 Main St"
	email := "john@example.com"
	subtotal := "0"
	tax := "0"
	paymentMethod := "Cash"
	deliveryMethod := "Dine in"
	addedCost := "0"
	status := "waiting"
	grandTotal := "0"
	orderID := 1
	productImage := "image.jpg"
	productName := "Coffee"
	price := 0
	size := "R"
	cost := 0
	qty := 1

	mockOrder := models.GetOrder{
			Orderlist_id:   1,
			Orderlist_uuid: "uuid1",
			First_name:      &firstName,
			Last_name:       &lastName,
			Phone:           &phone,
			Address:         &address,
			Email:           &email,
			Subtotal: 	     &subtotal,
			Tax: 		     &tax,
			Payment_method:  &paymentMethod,
			Delivery_method: &deliveryMethod,
			Added_cost:	     &addedCost,
			Status:			 &status,
			Grand_total:     &grandTotal,
			Products: []models.GetOrderDetail{
				{
					Order_id: &orderID,
					Product_image: &productImage,
					Product_name:    &productName,
					Price: &price,
					Size: &size,
					Added_cost: &cost,
					Qty: &qty,
				},
			},
		}

	orderRepositoryMock.On("GetHistoryOrder", mock.Anything).Return(&mockOrder, nil)
	orderDetailRepositoryMock.On("GetOrderDetails", mock.Anything).Return(&mockOrder.Products, nil)
	router.GET("/order/history", func(ctx *gin.Context) {
		ctx.Set("user_id", "1")
		handler.GetOrderHistory(ctx)
	})

	req, err := http.NewRequest("GET", "/order/history", nil)
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

func TestPatchOrder(t *testing.T) {
	router := gin.Default()
	orderRepositoryMock := new(repository.MockOrder)
	orderDetailRepositoryMock := new(repository.MockOrderDetails)
	handler := NewOrder(orderRepositoryMock, orderDetailRepositoryMock)

	orderRepositoryMock.On("UpdateOrder", mock.Anything).Return("Data updated", nil)
	router.PATCH("/order/:uuid", handler.PatchOrder)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"status": "waiting",
	})

	req, err := http.NewRequest("PATCH", "/order/uuid1", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code does not match")
	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "An error occurred when getting a response")

	assert.Equal(t, http.StatusOK, actualResponse.Status, "Status code does not match")
	assert.Equal(t, "Update data success", actualResponse.Message, "Response message does not match")
	assert.Equal(t, "Data updated", actualResponse.Data, "Data does not match")
}