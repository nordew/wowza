package service

import (
	"context"
	"strconv"
	"time"
	"wowza/internal/converter"
	"wowza/internal/dto"
	"wowza/internal/entity"
	"wowza/internal/storage"

	"go.uber.org/zap"
)

const (
	defaultFeedLimit = 10
	maxFeedLimit     = 50
)

type FeedService struct {
	postStorage storage.Post
	logger      *zap.Logger
}

func NewFeedService(deps Dependencies) *FeedService {
	return &FeedService{
		postStorage: deps.Storages.Post,
		logger:      deps.Logger,
	}
}

func (s *FeedService) GetFeed(ctx context.Context, cursor string, limit int) (*dto.FeedResponse, error) {
	parsedCursor, err := parseCursor(cursor)
	if err != nil {
		s.logger.Error("failed to parse feed cursor", zap.Error(err))
		return nil, err
	}

	normalizedLimit := normalizeLimit(limit)

	posts, err := s.postStorage.GetForFeed(ctx, parsedCursor, normalizedLimit)
	if err != nil {
		s.logger.Error("failed to get posts for feed", zap.Error(err))
		return nil, err
	}

	postDTOs := converter.ToPostResponseList(posts)
	nextCursor := calculateNextCursor(posts, normalizedLimit)

	return &dto.FeedResponse{
		Posts:      postDTOs,
		NextCursor: nextCursor,
	}, nil
}

func parseCursor(cursorStr string) (time.Time, error) {
	if cursorStr == "" {
		return time.Time{}, nil
	}

	unix, err := strconv.ParseInt(cursorStr, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(unix, 0), nil
}

func normalizeLimit(limit int) int {
	if limit <= 0 {
		return defaultFeedLimit
	}
	if limit > maxFeedLimit {
		return maxFeedLimit
	}
	
	return limit
}

func calculateNextCursor(posts []entity.Post, limit int) string {
	if len(posts) == limit && len(posts) > 0 {
		return strconv.FormatInt(posts[len(posts)-1].CreatedAt.Unix(), 10)
	}

	return ""
} 