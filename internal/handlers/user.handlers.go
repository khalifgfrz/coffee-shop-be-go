package handlers

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repository.RepoUser
}

func NewUser(r *repository.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostUser(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.CreateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerUser) FetchAll(ctx *gin.Context) {
	data, err := h.GetAllUser()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}