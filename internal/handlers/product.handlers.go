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

type HandlerProduct struct {
	repository.ProductRepositoryInterface
	pkg.CloudinaryInterface
}

func NewProduct(r repository.ProductRepositoryInterface, cld pkg.CloudinaryInterface) *HandlerProduct {
	return &HandlerProduct{r, cld}
}

func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	product := models.Product{}
	if err := ctx.ShouldBind(&product); err != nil {
		response.BadRequest("Create data failed", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&product)
	if err != nil {
		response.BadRequest("Create data failed", err.Error())
		return
	}

	result, err := h.CreateProduct(&product)
	if err != nil {
		response.BadRequest("Create data failed", err.Error())
		return
	}
	response.Created("Create data success", result)
}

func (h *HandlerProduct) GetProducts(ctx *gin.Context) {
	productName := ctx.Query("product_name")
	minPriceStr := ctx.Query("min_price")
	maxPriceStr := ctx.Query("max_price")
	category := ctx.Query("category")
	sortBy := ctx.Query("sort_by")
	pageStr := ctx.DefaultQuery("page","1")

	var minPrice, maxPrice, page int
	var err error

	if minPriceStr != "" {
		minPrice, err = strconv.Atoi(minPriceStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_price"})
			return
		}
	}

	if maxPriceStr != "" {
		maxPrice, err = strconv.Atoi(maxPriceStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid max_price"})
			return
		}
	}

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}
	}

	query := models.ProductQuery{
		Product_name: productName,
		MinPrice:    minPrice,
		MaxPrice:    maxPrice,
		Category:    category,
		SortBy:      sortBy,
		Page:        page,
	}

	response := pkg.NewResponse(ctx)
	result, err := h.GetAllProduct(&query)
	if err != nil {
		response.InternalServerError("Get data failed", err.Error())
		return
	}

	response.Success("Get data success", result)
}

func (h *HandlerProduct) GetProductDetail(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.GetDetailProduct(id)
	if err != nil {
		response.InternalServerError("Get data failed", err.Error())
		return
	}
	response.Success("Get data success", result)
}

func (h *HandlerProduct) ProductDelete(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.DeleteProduct(id)
	if err != nil {
		response.BadRequest("Delete data failed", err.Error())
		return
	}
	response.Success("Delete data success", result)
}

func (h *HandlerProduct) PatchProduct(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	body := models.Product{}
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
		response.BadRequest("Create data failed, upload file failed, wrong file type", err)
		return
	}
	randomNumber := rand.Int()
	fileName := fmt.Sprintf("coffee-product-%d", randomNumber)
	uploadResult, err := h.UploadFile(ctx, file, fileName)
	if err != nil {
		response.BadRequest("Create data failed, upload file failed", err.Error())
		return
	}
	body.Product_image = uploadResult.SecureURL
	result, err := h.UpdateProduct(&body,id)
	if err != nil {
		response.BadRequest("Update data failed", err.Error())
		return
	}
	response.Success("Update data success", result)
}