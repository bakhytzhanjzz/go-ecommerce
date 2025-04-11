package usecase

import (
	"context"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"
	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type categoryUseCase struct {
	categoryRepo repository.CategoryRepository
	log          *logrus.Logger
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository, log *logrus.Logger) CategoryUseCase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
		log:          log,
	}
}

func (uc *categoryUseCase) CreateCategory(ctx context.Context, category *entity.Category) error {
	return uc.categoryRepo.Create(ctx, category)
}

func (uc *categoryUseCase) GetCategory(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	return uc.categoryRepo.FindByID(ctx, id)
}

func (uc *categoryUseCase) UpdateCategory(ctx context.Context, category *entity.Category) error {
	return uc.categoryRepo.Update(ctx, category)
}

func (uc *categoryUseCase) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return uc.categoryRepo.Delete(ctx, id)
}

func (uc *categoryUseCase) ListCategories(ctx context.Context) ([]*entity.Category, error) {
	return uc.categoryRepo.List(ctx)
}
