package repository

import (
	"context"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, category *entity.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	var category entity.Category
	if err := r.db.WithContext(ctx).First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *entity.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *categoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Category{}, "id = ?", id).Error
}

func (r *categoryRepository) List(ctx context.Context) ([]*entity.Category, error) {
	var categories []*entity.Category
	if err := r.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
