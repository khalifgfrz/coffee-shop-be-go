package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func favorite(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/favorite")

	repo := repository.NewFavorite(d)
	handler := handlers.NewFavorite(repo)

	route.GET("/", handler.GetFavorites)
	route.GET("/:id", handler.GetFavoriteDetail)
	route.POST("/", handler.PostFavorite)
	route.DELETE("/:id", handler.FavoriteDelete)
	// route.PATCH("/:id", handler.FavoriteUpdate)
}