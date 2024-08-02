package handlers

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"net/http"
	"strconv"

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

	if err := h.CreateUser(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *HandlerUser) GetUsers(ctx *gin.Context) {
	data, err := h.GetAllUser()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *HandlerUser) GetUserDetail(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	data, err := h.GetDetailUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *HandlerUser) UserDelete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *HandlerUser) UserUpdate(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.UpdateUser(id,&user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, response)
}
