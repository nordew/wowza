package service

import (
	"context"
	"strconv"
	"time"
	"wowza/internal/converter"
	"wowza/internal/dto"
	"wowza/internal/entity"

	"go.uber.org/zap"
)

const (
	defaultFeedLimit = 10
	maxFeedLimit     = 50
)

func (s *Service) GetFeed(ctx context.Context, req dto.GetFeedRequest) (dto.GetFeedResponse, error) {
	cursor, err := parseCursor(req.CursorStr)
	if err != nil {
		s.logger.Error("failed to parse feed cursor", zap.Error(err))
		return dto.GetFeedResponse{}, err
	}

	limit := normalizeLimit(req.Limit)

	posts, err := s.postStorage.GetForFeed(ctx, cursor, limit)
	if err != nil {
		s.logger.Error("failed to get posts for feed", zap.Error(err))
		return dto.GetFeedResponse{}, err
	}

	postDTOs := converter.ToPostResponseList(posts)
	nextCursor := calculateNextCursor(posts, limit)

	return dto.GetFeedResponse{
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