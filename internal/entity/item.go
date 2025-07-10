package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
)

type Item struct {
	ID          string
	BusinessID  string
	Name        string
	Description string
	Price       float64
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewItem(
	id,
	businessID,
	name,
	description,
	imageURL string,
	price float64,
) (*Item, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid item id")
	}

	if _, err := uuid.Parse(businessID); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid business id")
	}

	if name == "" {
		return nil, errx.NewInternal().WithDescription("item name is required")
	}

	if price < 0 {
		return nil, errx.NewInternal().WithDescription("price cannot be negative")
	}

	now := time.Now()

	return &Item{
		ID:          id,
		BusinessID:  businessID,
		Name:        name,
		Description: description,
		Price:       price,
		ImageURL:    imageURL,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
} 