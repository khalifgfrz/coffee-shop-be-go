package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/middleware"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	var repo repository.UserRepositoryInterface = repository.NewUser(d)
	var cld pkg.Cloudinary = *pkg.NewCloudinaryUtil()
	handler := handlers.NewUser(repo, cld)

	route.GET("/", middleware.AuthJwtMiddleware(), handler.GetUsers)
	route.GET("/:id", middleware.AuthJwtMiddleware(), handler.GetUserDetail)
	route.POST("/", middleware.AuthJwtMiddleware(), handler.PostUser)
	route.PATCH("/:id", middleware.AuthJwtMiddleware(), handler.PatchUser)
	route.DELETE("/:id", middleware.AuthJwtMiddleware(), handler.UserDelete)
}