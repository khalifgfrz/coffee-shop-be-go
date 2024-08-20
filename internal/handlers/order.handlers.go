package handlers

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerOrder struct {
	repository.OrderRepositoryInterface
	repository.OrderDetailsRepositoryInterface
}

func NewOrder(orderRepo repository.OrderRepositoryInterface, orderDetailsRepo repository.OrderDetailsRepositoryInterface) *HandlerOrder {
	return &HandlerOrder{orderRepo,orderDetailsRepo}
}

func (h *HandlerOrder) PostOrder(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.Order{}
	
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.BadRequest("Invalid request", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("Create data failed", err.Error())
		return
	}

	orderID, err := h.CreateOrder(&body)
	if err != nil {
		response.InternalServerError("Failed to create order", err.Error())
		return
	}

	result, err := h.CreateOrderDetails(orderID, body.Products)
	if err != nil {
		response.InternalServerError("Failed to create order details", err.Error())
		return
	}

	response.Created("Create data success", result)
}

func (h *HandlerOrder) GetOrders(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page","1")
	response := pkg.NewResponse(ctx)

	var page int
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}
	}

	query := models.OrderQuery{
		Page: page,
	}
	orders, err := h.GetAllOrder(&query)
	if err != nil {
		response.InternalServerError("Get data failed", err.Error())
		return
	}

	for i := range *orders {
		order := &(*orders)[i]
		orderDetails, err := h.GetOrderDetails(order.Orderlist_id)
		if err != nil {
			response.InternalServerError("Get data failed", err.Error())
			return
		}
		order.Products = *orderDetails
	}

	response.Success("Get data success", orders)
}

func (h *HandlerOrder) GetOrderDetail(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	uuid := ctx.Param("uuid")
	order, err := h.GetDetailOrder(uuid)
	if err != nil {
		response.InternalServerError("Get data failed", err.Error())
		return
	}

	orderID := order.Orderlist_id
	orderDetails, err := h.GetOrderDetails(orderID)
	if err != nil {
		response.InternalServerError("Get data failed", err.Error())
		return
	}
	order.Products = *orderDetails
	
	response.Success("Get data success", order)
}

func (h *HandlerOrder) GetOrderHistory(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.InternalServerError("User ID not found", nil)
		return
	}
	id := userID.(string)
	order, err := h.GetHistoryOrder(id)
	if err != nil {
		response.InternalServerError("Get data failed", err.Error())
		return
	}

	orderID := order.Orderlist_id
	orderDetails, err := h.GetOrderDetails(orderID)
	if err != nil {
		response.InternalServerError("Get data failed", err.Error())
		return
	}
	order.Products = *orderDetails
	
	response.Success("Get data success", order)
}

func (h *HandlerOrder) PatchOrder(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	uuid := ctx.Param("uuid")
	body := models.Order{}
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("Update data failed", err.Error())
		return
	}
	result, err := h.UpdateOrder(&body,uuid)
	if err != nil {
		response.BadRequest("Update data failed", err.Error())
		return
	}
	response.Success("Update data success", result)
}

