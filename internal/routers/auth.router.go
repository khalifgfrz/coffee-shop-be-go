package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/auth")

	var repo repository.AuthRepositoryInterface = repository.NewAuth(d)
	handler := handlers.NewAuth(repo)

	route.POST("/register", handler.Register)
	route.POST("/login", handler.Login)
}