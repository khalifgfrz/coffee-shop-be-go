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

type HandlerUser struct {
	repository.UserRepositoryInterface
}

func NewUser(r repository.UserRepositoryInterface) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostUser(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	user := models.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&user)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	result, err := h.CreateUser(&user)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	response.Created("create data success", result)
}

func (h *HandlerUser) GetUsers(ctx *gin.Context) {
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

	query := models.UserQuery{
		Page: page,
	}

	response := pkg.NewResponse(ctx)
	result, err := h.GetAllUser(&query)
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}

func (h *HandlerUser) GetUserDetail(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.GetDetailUser(id)
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}
	response.Success("get data success", result)
}

func (h *HandlerUser) UserDelete(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.DeleteUser(id)
	if err != nil {
		response.BadRequest("delete data failed", err.Error())
		return
	}
	response.Success("delete data success", result)
}

func (h *HandlerUser) PatchUser(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	body := models.User{}
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}
	result, err := h.UpdateUser(&body,id)
	if err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}
	response.Success("update data success", result)
}
