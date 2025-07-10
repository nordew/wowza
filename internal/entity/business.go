package entity

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
)

type Business struct {
	ID          string
	UserID      string
	Name        string
	Description string
	WebsiteURL  string
	Location    string
	CategoryID  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Category    Category `gorm:"foreignKey:CategoryID"`
	User        User     `gorm:"foreignKey:UserID"`
}

func NewBusiness(
	id,
	userID,
	name,
	description,
	websiteURL,
	location,
	categoryID string,
) (*Business, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid business id")
	}

	if _, err := uuid.Parse(userID); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid user id")
	}

	if _, err := uuid.Parse(categoryID); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid category id")
	}

	if name == "" {
		return nil, errx.NewInternal().WithDescription("business name is required")
	}

	if websiteURL != "" {
		if _, err := url.ParseRequestURI(websiteURL); err != nil {
			return nil, errx.NewInternal().WithDescription("invalid website url")
		}
	}

	now := time.Now()

	return &Business{
		ID:          id,
		UserID:      userID,
		Name:        name,
		Description: description,
		WebsiteURL:  websiteURL,
		Location:    location,
		CategoryID:  categoryID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
} 