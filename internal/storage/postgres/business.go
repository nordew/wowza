package postgres

import (
	"context"
	"errors"
	"wowza/internal/entity"

	"github.com/nordew/go-errx"
	"gorm.io/gorm"
)

type BusinessStorage struct {
	db *gorm.DB
}

func NewBusinessStorage(db *gorm.DB) *BusinessStorage {
	return &BusinessStorage{db: db}
}

func (s *BusinessStorage) Create(ctx context.Context, business *entity.Business, _ []string) error {
	if err := s.db.WithContext(ctx).Create(business).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to create business", err)
	}
	return nil
}

func (s *BusinessStorage) GetByID(ctx context.Context, id string) (*entity.Business, error) {
	var business entity.Business
	if err := s.db.WithContext(ctx).Preload("Category").First(&business, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewNotFound().WithDescription("business not found")
		}
		return nil, errx.NewInternal().WithDescriptionAndCause("failed to get business by id", err)
	}
	return &business, nil
}

func (s *BusinessStorage) Update(ctx context.Context, business *entity.Business, _ []string) error {
	if err := s.db.WithContext(ctx).Save(business).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to update business", err)
	}
	return nil
}

func (s *BusinessStorage) Delete(ctx context.Context, id string) error {
	if err := s.db.WithContext(ctx).Delete(&entity.Business{}, "id = ?", id).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to delete business", err)
	}
	return nil
}

func (s *BusinessStorage) GetByUserID(ctx context.Context, userID string) ([]entity.Business, error) {
	var businesses []entity.Business
	if err := s.db.WithContext(ctx).Preload("Category").Where("user_id = ?", userID).Find(&businesses).Error; err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause("failed to get businesses by user id", err)
	}
	return businesses, nil
} 