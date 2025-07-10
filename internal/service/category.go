package service

import (
	"context"
	"wowza/internal/converter"
	"wowza/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := s.categoryStorage.GetAll(ctx)
	if err != nil {
		s.logger.Error("failed to get all categories", zap.Error(err))
		return nil, err
	}

	return converter.ToCategoryResponseList(categories), nil
} 