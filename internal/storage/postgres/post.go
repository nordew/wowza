package postgres

import (
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

func (s *PostStorage) Create(post *entity.Post) error {
	if err := s.db.Create(post).Error; err != nil {
		return errx.NewInternal().WithDescription("failed to create post")
	}

	return nil
}
