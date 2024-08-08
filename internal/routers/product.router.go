package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/middleware"
	"khalifgfrz/coffee-shop-be-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	var repo repository.ProductRepositoryInterface = repository.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.GET("/", handler.GetProducts)
	route.GET("/:id", handler.GetProductDetail)
	route.POST("/", middleware.AuthJwtMiddleware(), handler.PostProduct)
	route.DELETE("/:id", middleware.AuthJwtMiddleware(), handler.ProductDelete)
	route.PATCH("/:id", middleware.AuthJwtMiddleware(), handler.PatchProduct)
}