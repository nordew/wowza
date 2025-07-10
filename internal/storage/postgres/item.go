package postgres

import (
	"context"
	"errors"
	"wowza/internal/entity"

	"github.com/nordew/go-errx"
	"gorm.io/gorm"
)

type ItemStorage struct {
	db *gorm.DB
}

func NewItemStorage(db *gorm.DB) *ItemStorage {
	return &ItemStorage{db: db}
}

func (s *ItemStorage) Create(ctx context.Context, item *entity.Item) error {
	if err := s.db.WithContext(ctx).Create(item).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to create item", err)
	}
	return nil
}

func (s *ItemStorage) GetByID(ctx context.Context, id string) (*entity.Item, error) {
	var item entity.Item
	if err := s.db.WithContext(ctx).Preload("Reviews").First(&item, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewNotFound().WithDescription("item not found")
		}
		return nil, errx.NewInternal().WithDescriptionAndCause("failed to get item by id", err)
	}
	return &item, nil
}

func (s *ItemStorage) Update(ctx context.Context, item *entity.Item) error {
	if err := s.db.WithContext(ctx).Save(item).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to update item", err)
	}
	return nil
}

func (s *ItemStorage) Delete(ctx context.Context, id string) error {
	if err := s.db.WithContext(ctx).Delete(&entity.Item{}, "id = ?", id).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to delete item", err)
	}
	return nil
}

func (s *ItemStorage) GetByBusinessID(ctx context.Context, businessID string) ([]entity.Item, error) {
	var items []entity.Item
	if err := s.db.WithContext(ctx).Where("business_id = ?", businessID).Find(&items).Error; err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause("failed to get items by business id", err)
	}
	return items, nil
} 