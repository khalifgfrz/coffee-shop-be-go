package handlers

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type HandlerOrder struct {
	db *sqlx.DB
	repository.OrderRepositoryInterface
	repository.OrderDetailsRepositoryInterface
}

func NewOrder(r repository.OrderRepositoryInterface, rd repository.OrderDetailsRepositoryInterface) *HandlerOrder {
	return &HandlerOrder{
		OrderRepositoryInterface: r,
		OrderDetailsRepositoryInterface: rd,
	}
}

func (h *HandlerOrder) CreateNewOrder(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	var body models.OrderDetailsBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.BadRequest("Invalid request", err.Error())
		return
	}

	tx, err := h.db.Beginx()
	if err != nil {
		response.InternalServerError("Database connection error", err.Error())
		return
	}
	defer tx.Rollback()

	orderID, err := h.CreateOrder(&body.Order, tx)
	if err != nil {
		response.InternalServerError("Failed to create order", err.Error())
		return
	}

	detailResult, err := h.CreateOrderDetails(orderID, body.Products)
	if err != nil {
		response.InternalServerError("Failed to create order details", err.Error())
		return
	}

	if err := tx.Commit(); err != nil {
		response.InternalServerError("Failed to commit transaction", err.Error())
		return
	}

	response.Created("Success", gin.H{
		"orderID": orderID,
		"details": detailResult,
	})
}