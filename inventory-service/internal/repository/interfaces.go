package repository

import (
	"context"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"
	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, limit int, filters map[string]interface{}) ([]*entity.Product, error)
}

type CategoryRepository interface {
	Create(ctx context.Context, category *entity.Category) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]*entity.Category, error)
}
