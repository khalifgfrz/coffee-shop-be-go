package handlers

import (
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	repository.UserRepositoryInterface
	pkg.CloudinaryInterface
}

func NewUser(r repository.UserRepositoryInterface, cld pkg.CloudinaryInterface) *HandlerUser {
	return &HandlerUser{r, cld}
}

func (h *HandlerUser) PostUser(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	user := models.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		response.BadRequest("Create data failed", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&user)
	if err != nil {
		response.BadRequest("Create data failed", err.Error())
		return
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		response.BadRequest("Create data failed", err.Error())
		return
	}

	result, err := h.CreateUser(&user)
	if err != nil {
		response.BadRequest("Create data failed", err.Error())
		return
	}
	response.Created("Create data success", result)
}

func (h *HandlerUser) GetUsers(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page","1")
	response := pkg.NewResponse(ctx)

	var page int
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			response.BadRequest("Invalid page number", err.Error())
			return
		}
	}

	query := models.UserQuery{
		Page: page,
	}

	result, err := h.GetAllUser(&query)
	if err != nil {
		response.InternalServerError("Get data failed", err.Error())
		return
	}

	response.Success("Get data success", result)
}

func (h *HandlerUser) GetUserDetail(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.InternalServerError("User ID not found", nil)
		return
	}
	id := userID.(string)
	result, err := h.GetDetailUser(id)
	if err != nil {
		response.InternalServerError("Get data failed", err.Error())
		return
	}
	response.Success("Get data success", result)
}

func (h *HandlerUser) UserDelete(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.InternalServerError("User ID not found", nil)
		return
	}
	id := userID.(string)
	result, err := h.DeleteUser(id)
	if err != nil {
		response.BadRequest("Delete data failed", err.Error())
		return
	}
	response.Success("Delete data success", result)
}

func (h *HandlerUser) PatchUser(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.InternalServerError("User ID not found", nil)
		return
	}
	id := userID.(string)
	body := models.User{}
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("Update data failed", err.Error())
		return
	}
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		response.BadRequest("Create data failed, upload file failed", err.Error())
		return
	}
	if header.Size > 1*1024*1024 {
		response.BadRequest("Create data failed, upload file failed, file too large", nil)
		return
	}
	mimeType := header.Header.Get("Content-Type")
	if mimeType != "image/jpg" && mimeType != "image/png" {
		response.BadRequest("Create data failed, upload file failed, wrong file type", nil)
		return
	}
	randomNumber := rand.Int()
	fileName := fmt.Sprintf("coffee-user-%d", randomNumber)
	uploadResult, err := h.UploadFile(ctx, file, fileName)
	if err != nil {
		response.BadRequest("Create data failed, upload file failed", err.Error())
		return
	}
	body.User_image = uploadResult.SecureURL
	result, err := h.UpdateUser(&body,id)
	if err != nil {
		response.BadRequest("Update data failed", err.Error())
		return
	}
	response.Success("Update data success", result)
}
