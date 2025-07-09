package entity

import (
	"time"
	"wowza/internal/dto"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
)

type Post struct {
	ID            string
	UserID        string
	VideoURL      string
	Description   string
	Duration      float64
	Visibility    dto.PostVisibility
	Hashtags      []string
	Tags          []string
	LikesCount    int64
	CommentsCount int64
	ViewsCount    int64
	AllowComments bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewPost(id string, req *dto.CreatePostRequest, videoURL string) (*Post, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid id")
	}

	if _, err := uuid.Parse(req.UserID); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid user id")
	}

	if videoURL == "" {
		return nil, errx.NewInternal().WithDescription("video url is required")
	}

	if req.Duration <= 0 {
		return nil, errx.NewInternal().WithDescription("duration must be positive")
	}

	if !req.Visibility.IsValid() {
		return nil, errx.NewInternal().WithDescription("invalid visibility")
	}

	now := time.Now()

	return &Post{
		ID:            id,
		UserID:        req.UserID,
		VideoURL:      videoURL,
		Description:   req.Description,
		Duration:      req.Duration,
		Visibility:    req.Visibility,
		Hashtags:      req.Hashtags,
		Tags:          req.Tags,
		LikesCount:    0,
		CommentsCount: 0,
		ViewsCount:    0,
		AllowComments: req.AllowComments,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}
 