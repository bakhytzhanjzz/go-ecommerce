package usecase

import (
	"context"
	"errors"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"
	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type productUseCase struct {
	productRepo repository.ProductRepository
	log         *logrus.Logger
}

func NewProductUsecase(productRepo repository.ProductRepository, log *logrus.Logger) ProductUseCase {
	return &productUseCase{
		productRepo: productRepo,
		log:         log,
	}
}

func (uc *productUseCase) CreateProduct(ctx context.Context, product *entity.Product) error {
	if err := uc.productRepo.Create(ctx, product); err != nil {
		uc.log.WithError(err).Error("Failed to create product")
		return err
	}
	return nil
}

func (uc *productUseCase) GetProduct(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	product, err := uc.productRepo.FindByID(ctx, id)
	if err != nil {
		uc.log.WithError(err).Error("Failed to get product")
		return nil, err
	}
	return product, nil
}

func (uc *productUseCase) UpdateProduct(ctx context.Context, product *entity.Product) error {
	if err := uc.productRepo.Update(ctx, product); err != nil {
		uc.log.WithError(err).Error("Failed to update product")
		return err
	}
	return nil
}

func (uc *productUseCase) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if err := uc.productRepo.Delete(ctx, id); err != nil {
		uc.log.WithError(err).Error("Failed to delete product")
		return err
	}
	return nil
}

func (uc *productUseCase) ListProducts(ctx context.Context, page, limit int, filters map[string]interface{}) ([]*entity.Product, error) {
	products, err := uc.productRepo.List(ctx, page, limit, filters)
	if err != nil {
		uc.log.WithError(err).Error("Failed to list products")
		return nil, err
	}
	return products, nil
}

func (uc *productUseCase) DecreaseStock(ctx context.Context, productID uuid.UUID, amount int) error {
	product, err := uc.productRepo.FindByID(ctx, productID)
	if err != nil {
		uc.log.WithError(err).Error("Failed to find product for stock decrease")
		return err
	}

	if product.Stock < amount {
		return errors.New("insufficient stock")
	}

	product.Stock -= amount
	return uc.productRepo.Update(ctx, product)
}

func (uc *productUseCase) IncreaseStock(ctx context.Context, productID uuid.UUID, amount int) error {
	product, err := uc.productRepo.FindByID(ctx, productID)
	if err != nil {
		uc.log.WithError(err).Error("Failed to find product for stock increase")
		return err
	}

	product.Stock += amount
	return uc.productRepo.Update(ctx, product)
}
