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
	var cld pkg.CloudinaryInterface = pkg.NewCloudinaryUtil()
	handler := handlers.NewUser(repo, cld)

	route.GET("/datauser", middleware.Auth("admin"), handler.GetUsers)
	route.GET("/profile", middleware.Auth("admin", "user"), handler.GetUserDetail)
	route.POST("/", middleware.Auth("admin"), handler.PostUser)
	route.PATCH("/settings", middleware.Auth("admin", "user"), handler.PatchUser)
	route.DELETE("/delete", middleware.Auth("user"), handler.UserDelete)
}