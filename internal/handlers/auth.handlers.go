package handlers

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerAuth struct {
	repository.AuthRepositoryInterface
}

func NewAuth(r repository.AuthRepositoryInterface) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Register(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.Auth{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	_, err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	body.Password, err = pkg.HashPassword(body.Password)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	result, err := h.RegisterUser(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	response.Created("create data success", result)
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.Auth{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("login failed", err.Error())
		return
	}
	_, err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("login failed", err.Error())
		return
	}

	result, err := h.LoginUser(body.Email)
	if err != nil {
		response.BadRequest("login failed", err.Error())
		return
	}

	err = pkg.VerifyPassword(result.Password, body.Password)
	if err != nil {
		response.Unauthorized("wrong password", err.Error())
		return
	}

	jwt := pkg.NewJWT(result.User_id, result.Email, result.Role)
	token, err := jwt.GenerateToken()
	if err != nil {
		response.Unauthorized("failed generate token", err.Error())
		return
	}

	response.Created("login success", token)
}