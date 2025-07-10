package service

import (
	"context"
	"wowza/internal/dto"
	"wowza/internal/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *Service) CreateItem(ctx context.Context, req dto.CreateItemRequest) (*dto.ItemResponse, error) {
	item, err := entity.NewItem(
		uuid.NewString(),
		req.BusinessID,
		req.Name,
		req.Description,
		req.ImageURL,
		req.Price,
	)
	if err != nil {
		s.logger.Error("failed to create new item entity", zap.Error(err))
		return nil, err
	}

	if err := s.itemStorage.Create(ctx, item); err != nil {
		s.logger.Error("failed to create item in storage", zap.Error(err))
		return nil, err
	}

	return s.GetItemByID(ctx, item.ID)
}

func (s *Service) GetItemByID(ctx context.Context, id string) (*dto.ItemResponse, error) {
	item, err := s.itemStorage.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get item by id", zap.Error(err))
		return nil, err
	}

	return &dto.ItemResponse{
		ID:          item.ID,
		BusinessID:  item.BusinessID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		ImageURL:    item.ImageURL,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}, nil
}

func (s *Service) UpdateItem(ctx context.Context, id string, req dto.UpdateItemRequest) (*dto.ItemResponse, error) {
	item, err := s.itemStorage.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		item.Name = *req.Name
	}
	if req.Description != nil {
		item.Description = *req.Description
	}
	if req.Price != nil {
		item.Price = *req.Price
	}
	if req.ImageURL != nil {
		item.ImageURL = *req.ImageURL
	}

	if err := s.itemStorage.Update(ctx, item); err != nil {
		return nil, err
	}

	return s.GetItemByID(ctx, id)
}

func (s *Service) DeleteItem(ctx context.Context, id string) error {
	return s.itemStorage.Delete(ctx, id)
}

func (s *Service) GetItemsByBusinessID(ctx context.Context, businessID string) ([]dto.ItemResponse, error) {
	items, err := s.itemStorage.GetByBusinessID(ctx, businessID)
	if err != nil {
		s.logger.Error("failed to get items by business id", zap.Error(err))
		return nil, err
	}

	res := make([]dto.ItemResponse, len(items))
	for i, item := range items {
		res[i] = dto.ItemResponse{
			ID:          item.ID,
			BusinessID:  item.BusinessID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			ImageURL:    item.ImageURL,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}

	return res, nil
} 