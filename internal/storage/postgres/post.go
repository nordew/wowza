package postgres

import (
	"context"
	"wowza/internal/entity"

	"github.com/nordew/go-errx"
	"gorm.io/gorm"
)

type PostStorage struct {
	db *gorm.DB
}

func NewPostStorage(db *gorm.DB) *PostStorage {
	return &PostStorage{db: db}
}

func (s *PostStorage) Create(ctx context.Context, post *entity.Post) error {
	if err := s.db.WithContext(ctx).Create(post).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to create post", err)
	}

	return nil
}
