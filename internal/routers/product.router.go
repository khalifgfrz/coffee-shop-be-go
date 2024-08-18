package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/middleware"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	var repo repository.ProductRepositoryInterface = repository.NewProduct(d)
	var cld pkg.Cloudinary = *pkg.NewCloudinaryUtil()
	handler := handlers.NewProduct(repo, cld)

	route.GET("/", handler.GetProducts)
	route.GET("/:id", handler.GetProductDetail)
	route.POST("/", middleware.Auth("admin"), handler.PostProduct)
	route.DELETE("/:id", middleware.Auth("admin"), handler.ProductDelete)
	route.PATCH("/:id", middleware.Auth("admin"), handler.PatchProduct)
}