package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
)

type Category struct {
	ID         string     
	Name       string    
	CreatedAt  time.Time  
	UpdatedAt  time.Time  
}

func NewCategory(id, name string) (*Category, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid id")
	}

	if name == "" {
		return nil, errx.NewInternal().WithDescription("category name is required")
	}

	now := time.Now()

	return &Category{
		ID:        id,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
} 