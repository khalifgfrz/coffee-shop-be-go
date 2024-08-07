package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repository.NewUser(d)
	handler := handlers.NewUser(repo)

	route.GET("/", handler.GetUsers)
	route.GET("/:id", handler.GetUserDetail)
	route.POST("/", handler.PostUser)
	route.PATCH("/:id", handler.PatchUser)
	route.DELETE("/:id", handler.UserDelete)
}