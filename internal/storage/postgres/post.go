package postgres

import (
	"context"
	"time"
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

func (s *PostStorage) GetForFeed(ctx context.Context, cursor time.Time, limit int) ([]entity.Post, error) {
	var posts []entity.Post
	query := s.db.WithContext(ctx).Order("created_at desc").Limit(limit)

	if !cursor.IsZero() {
		query = query.Where("created_at < ?", cursor)
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause("failed to get posts for feed", err)
	}

	return posts, nil
}
