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

type HandlerProduct struct {
	repository.ProductRepositoryInterface
}

func NewProduct(r repository.ProductRepositoryInterface) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	product := models.Product{}
	if err := ctx.ShouldBind(&product); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&product)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	result, err := h.CreateProduct(&product)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	response.Created("create data success", result)
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
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}

func (h *HandlerProduct) GetProductDetail(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.GetDetailProduct(id)
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}
	response.Success("get data success", result)
}

func (h *HandlerProduct) ProductDelete(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.DeleteProduct(id)
	if err != nil {
		response.BadRequest("delete data failed", err.Error())
		return
	}
	response.Success("delete data success", result)
}

func (h *HandlerProduct) PatchProduct(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	body := models.Product{}
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}
	result, err := h.UpdateProduct(&body,id)
	if err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}
	response.Success("update data success", result)
}