package handlers

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerFavorite struct {
	*repository.RepoFavorite
}

func NewFavorite(r *repository.RepoFavorite) *HandlerFavorite {
	return &HandlerFavorite{r}
}

func (h *HandlerFavorite) PostFavorite(ctx *gin.Context) {
	favorite := models.PostFavorite{}

	if err := ctx.ShouldBind(&favorite); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.CreateFavorite(&favorite); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Favorite product added successfully"})
}

func (h *HandlerFavorite) GetFavorites(ctx *gin.Context) {
	pageStr := ctx.Query("page")

	var page int
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}
	} else {
		page = 1
	}

	query := models.FavoriteQuery{
		Page: page,
	}
	data, err := h.GetAllFavorite(&query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(*data) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No favorite found"})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *HandlerFavorite) GetFavoriteDetail(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid favorite ID"})
		return
	}

	data, err := h.GetDetailFavorite(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching favorite details"})
		return
	}

	if data == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Favorite not found"})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *HandlerFavorite) FavoriteDelete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid favorite ID"})
		return
	}

	if err := h.DeleteFavorite(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete favorite product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Favorite product deleted successfully"})
}

func (h *HandlerFavorite) PatchFavorite(ctx *gin.Context) {
	var favorite models.UpdateFavorite
	if err := ctx.ShouldBindJSON(&favorite); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid favorite ID"})
		return
	}

	updatedFavorite, err := h.UpdateFavorite(&favorite, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedFavorite)
}
