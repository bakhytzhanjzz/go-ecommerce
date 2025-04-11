package controller

import (
	"net/http"
	"strconv"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/usecase"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	productUsecase usecase.ProductUseCase
}

func NewProductController(productUsecase usecase.ProductUseCase) *ProductController {
	return &ProductController{
		productUsecase: productUsecase,
	}
}

type createProductRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"required"`
	Stock       int       `json:"stock" binding:"required"`
	CategoryID  uuid.UUID `json:"category_id"`
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := &entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
	}

	if err := c.productUsecase.CreateProduct(ctx.Request.Context(), product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) GetProduct(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	product, err := c.productUsecase.GetProduct(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type updateProductRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  uuid.UUID `json:"category_id"`
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	var req updateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.productUsecase.GetProduct(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	if req.Stock != 0 {
		product.Stock = req.Stock
	}
	if req.CategoryID != uuid.Nil {
		product.CategoryID = req.CategoryID
	}

	if err := c.productUsecase.UpdateProduct(ctx.Request.Context(), product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	if err := c.productUsecase.DeleteProduct(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *ProductController) ListProducts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	filters := make(map[string]interface{})
	if categoryID := ctx.Query("category_id"); categoryID != "" {
		if id, err := uuid.Parse(categoryID); err == nil {
			filters["category_id"] = id
		}
	}
	if name := ctx.Query("name"); name != "" {
		filters["name"] = name
	}

	products, err := c.productUsecase.ListProducts(ctx.Request.Context(), page, limit, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
