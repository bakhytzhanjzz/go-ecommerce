package usecase

import (
	"context"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"
	"github.com/google/uuid"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product *entity.Product) error
	GetProduct(ctx context.Context, id uuid.UUID) (*entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	ListProducts(ctx context.Context, page, limit int, filters map[string]interface{}) ([]*entity.Product, error)
	DecreaseStock(ctx context.Context, productID uuid.UUID, amount int) error
	IncreaseStock(ctx context.Context, productID uuid.UUID, amount int) error
}

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, category *entity.Category) error
	GetCategory(ctx context.Context, id uuid.UUID) (*entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id uuid.UUID) error
	ListCategories(ctx context.Context) ([]*entity.Category, error)
}
