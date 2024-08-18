package handlers

import (
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	repository.UserRepositoryInterface
	pkg.Cloudinary
}

func NewUser(r repository.UserRepositoryInterface, cld pkg.Cloudinary) *HandlerUser {
	return &HandlerUser{r, cld}
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
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.InternalServerError("User ID not found", nil)
		return
	}
	id := userID.(string)
	result, err := h.GetDetailUser(id)
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}
	response.Success("get data success", result)
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
		response.BadRequest("delete data failed", err.Error())
		return
	}
	response.Success("delete data success", result)
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
		response.BadRequest("update data failed", err.Error())
		return
	}
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		response.BadRequest("create data failed, upload file failed", err.Error())
		return
	}
	if header.Size > 1*1024*1024 {
		response.BadRequest("create data failed, upload file failed, file too large", nil)
		return
	}
	mimeType := header.Header.Get("Content-Type")
	if mimeType != "image/jpg" && mimeType != "image/png" {
		response.BadRequest("create data failed, upload file failed, wrong file type", nil)
		return
	}
	randomNumber := rand.Int()
	fileName := fmt.Sprintf("coffee-user-%d", randomNumber)
	uploadResult, err := h.UploadFile(ctx, file, fileName)
	if err != nil {
		response.BadRequest("create data failed, upload file failed", err.Error())
		return
	}
	body.User_image = uploadResult.SecureURL
	result, err := h.UpdateUser(&body,id)
	if err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}
	response.Success("update data success", result)
}
