package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/middleware"
	"khalifgfrz/coffee-shop-be-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func favorite(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/favorite")

	var repo repository.FavoriteRepositoryInterface = repository.NewFavorite(d)
	handler := handlers.NewFavorite(repo)

	route.GET("/", middleware.AuthJwtMiddleware(), handler.GetFavorites)
	route.GET("/:id", middleware.AuthJwtMiddleware(), handler.GetFavoriteDetail)
	route.POST("/", middleware.AuthJwtMiddleware(), handler.PostFavorite)
	route.DELETE("/:id", middleware.AuthJwtMiddleware(), handler.FavoriteDelete)
	route.PATCH("/:id", middleware.AuthJwtMiddleware(), handler.PatchFavorite)
}