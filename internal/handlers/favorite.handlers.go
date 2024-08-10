package handlers

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerFavorite struct {
	repository.FavoriteRepositoryInterface
}

func NewFavorite(r repository.FavoriteRepositoryInterface) *HandlerFavorite {
	return &HandlerFavorite{r}
}

func (h *HandlerFavorite) PostFavorite(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	favorite := models.Favorite{}
	if err := ctx.ShouldBind(&favorite); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&favorite)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	result, err := h.CreateFavorite(&favorite)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	response.Created("create data success", result)
}

func (h *HandlerFavorite) GetFavorites(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page","1")

	var page int
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}
	}

	query := models.FavoriteQuery{
		Page: page,
	}
	response := pkg.NewResponse(ctx)
	result, err := h.GetAllFavorite(&query)
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}

func (h *HandlerFavorite) GetFavoriteDetail(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.GetDetailFavorite(id)
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}
	response.Success("get data success", result)
}

func (h *HandlerFavorite) FavoriteDelete(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.DeleteFavorite(id)
	if err != nil {
		response.BadRequest("delete data failed", err.Error())
		return
	}
	response.Success("delete data success", result)
}

func (h *HandlerFavorite) PatchFavorite(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	body := models.Favorite{}
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}
	result, err := h.UpdateFavorite(&body,id)
	if err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}
	response.Success("update data success", result)
}
