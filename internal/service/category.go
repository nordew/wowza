package service

import (
	"context"
	"wowza/internal/converter"
	"wowza/internal/dto"
	"wowza/internal/storage"

	"go.uber.org/zap"
)

type CategoryService struct {
	categoryStorage storage.Category
	logger          *zap.Logger
}

func NewCategoryService(deps Dependencies) *CategoryService {
	return &CategoryService{
		categoryStorage: deps.Storages.Category,
		logger:          deps.Logger,
	}
}

func (s *CategoryService) GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := s.categoryStorage.GetAll(ctx)
	if err != nil {
		s.logger.Error("failed to get all categories", zap.Error(err))
		return nil, err
	}

	return converter.ToCategoryResponseList(categories), nil
} 