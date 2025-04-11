package repository

import (
	"context"
	"errors"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"
	"github.com/google/uuid"
)

type InMemoryProductRepository struct {
	products map[uuid.UUID]*entity.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[uuid.UUID]*entity.Product),
	}
}

func (r *InMemoryProductRepository) Create(ctx context.Context, product *entity.Product) error {
	if _, exists := r.products[product.ID]; exists {
		return errors.New("product already exists")
	}
	r.products[product.ID] = product
	return nil
}

func (r *InMemoryProductRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (r *InMemoryProductRepository) Update(ctx context.Context, product *entity.Product) error {
	if _, exists := r.products[product.ID]; !exists {
		return errors.New("product not found")
	}
	r.products[product.ID] = product
	return nil
}

func (r *InMemoryProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if _, exists := r.products[id]; !exists {
		return errors.New("product not found")
	}
	delete(r.products, id)
	return nil
}

func (r *InMemoryProductRepository) List(ctx context.Context, page, limit int, filters map[string]interface{}) ([]*entity.Product, error) {
	var products []*entity.Product
	for _, p := range r.products {
		products = append(products, p)
	}
	return products, nil
}
