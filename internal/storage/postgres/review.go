package postgres

import (
	"context"
	"errors"
	"wowza/internal/entity"

	"github.com/nordew/go-errx"
	"gorm.io/gorm"
)

type ReviewStorage struct {
	db *gorm.DB
}

func NewReviewStorage(db *gorm.DB) *ReviewStorage {
	return &ReviewStorage{db: db}
}

func (s *ReviewStorage) Create(ctx context.Context, review *entity.Review) error {
	if err := s.db.WithContext(ctx).Create(review).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to create review", err)
	}
	return nil
}

func (s *ReviewStorage) GetByID(ctx context.Context, id string) (*entity.Review, error) {
	var review entity.Review
	if err := s.db.WithContext(ctx).First(&review, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewNotFound().WithDescription("review not found")
		}
		return nil, errx.NewInternal().WithDescriptionAndCause("failed to get review by id", err)
	}
	return &review, nil
}

func (s *ReviewStorage) Update(ctx context.Context, review *entity.Review) error {
	if err := s.db.WithContext(ctx).Save(review).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to update review", err)
	}
	return nil
}

func (s *ReviewStorage) Delete(ctx context.Context, id string) error {
	if err := s.db.WithContext(ctx).Delete(&entity.Review{}, "id = ?", id).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to delete review", err)
	}
	return nil
}

func (s *ReviewStorage) GetByItemID(ctx context.Context, itemID string) ([]entity.Review, error) {
	var reviews []entity.Review
	if err := s.db.WithContext(ctx).Where("item_id = ?", itemID).Find(&reviews).Error; err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause("failed to get reviews by item id", err)
	}
	return reviews, nil
} 