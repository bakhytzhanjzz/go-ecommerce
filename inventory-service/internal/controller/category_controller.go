package controller

import (
	"net/http"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"
	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryController struct {
	categoryUsecase usecase.CategoryUseCase
}

func NewCategoryController(categoryUsecase usecase.CategoryUseCase) *CategoryController {
	return &CategoryController{
		categoryUsecase: categoryUsecase,
	}
}

func (c *CategoryController) Create(ctx *gin.Context) {
	var category entity.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.categoryUsecase.CreateCategory(ctx.Request.Context(), &category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

func (c *CategoryController) List(ctx *gin.Context) {
	categories, err := c.categoryUsecase.ListCategories(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) Get(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	category, err := c.categoryUsecase.GetCategory(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	var updateData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := c.categoryUsecase.GetCategory(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	if updateData.Name != "" {
		category.Name = updateData.Name
	}
	if updateData.Description != "" {
		category.Description = updateData.Description
	}

	if err := c.categoryUsecase.UpdateCategory(ctx.Request.Context(), category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	if err := c.categoryUsecase.DeleteCategory(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
