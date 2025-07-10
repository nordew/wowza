package postgres

import (
	"context"
	"wowza/internal/entity"

	"github.com/nordew/go-errx"
	"gorm.io/gorm"
)

type CategoryStorage struct {
	db *gorm.DB
}

func NewCategoryStorage(db *gorm.DB) *CategoryStorage {
	return &CategoryStorage{db: db}
}

func (s *CategoryStorage) GetAll(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category
	if err := s.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause("failed to get all categories", err)
	}
	return categories, nil
} 