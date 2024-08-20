package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func order(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/order")

	var orderRepo repository.OrderRepositoryInterface = repository.NewOrder(d)
	var orderDetailsRepo repository.OrderDetailsRepositoryInterface = repository.NewOrderDetails(d)
	handler := handlers.NewOrder(orderRepo,orderDetailsRepo)
	
	route.POST("/", handler.CreateNewOrder)
}