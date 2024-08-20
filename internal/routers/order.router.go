package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/middleware"
	"khalifgfrz/coffee-shop-be-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func order(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/order")

	var orderRepo repository.OrderRepositoryInterface = repository.NewOrder(d)
	var orderDetailsRepo repository.OrderDetailsRepositoryInterface = repository.NewOrderDetails(d)
	handler := handlers.NewOrder(orderRepo,orderDetailsRepo)
	
	route.POST("/", middleware.Auth("user"), handler.PostOrder)
	route.GET("/:uuid", middleware.Auth("user"), handler.GetOrderDetail)
	route.GET("/history", middleware.Auth("user"), handler.GetOrderHistory)
	route.GET("/", middleware.Auth("admin"), handler.GetOrders)
	route.PATCH("/:uuid", middleware.Auth("admin"), handler.PatchOrder)
}