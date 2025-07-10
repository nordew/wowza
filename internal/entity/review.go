package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
)

type Review struct {
	ID          string
	UserID      string
	ItemID      string
	Rating      int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewReview(
	id,
	userID,
	itemID,
	description string,
	rating int,
) (*Review, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid review id")
	}

	if _, err := uuid.Parse(userID); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid user id")
	}

	if _, err := uuid.Parse(itemID); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid item id")
	}

	if rating < 1 || rating > 5 {
		return nil, errx.NewInternal().WithDescription("rating must be between 1 and 5")
	}

	now := time.Now()

	return &Review{
		ID:          id,
		UserID:      userID,
		ItemID:      itemID,
		Rating:      rating,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
} 