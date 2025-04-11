package repository

import (
	"context"
	"fmt"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *entity.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.WithContext(ctx).Preload("Category").First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, product *entity.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Product{}, "id = ?", id).Error
}

func (r *productRepository) List(ctx context.Context, page, limit int, filters map[string]interface{}) ([]*entity.Product, error) {
	var products []*entity.Product

	query := r.db.WithContext(ctx).Preload("Category")

	// Apply filters
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Apply pagination
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
