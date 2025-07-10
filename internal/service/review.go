package service

import (
	"context"
	"wowza/internal/dto"
	"wowza/internal/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *Service) CreateReview(ctx context.Context, req dto.CreateReviewRequest) (*dto.ReviewResponse, error) {
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		s.logger.Error("failed to get user id from context")
		return nil, &serviceErr{msg: "unauthorized", code: 401}
	}

	review, err := entity.NewReview(
		uuid.NewString(),
		userID,
		req.ItemID,
		req.Description,
		req.Rating,
	)
	if err != nil {
		s.logger.Error("failed to create new review entity", zap.Error(err))
		return nil, err
	}

	if err := s.reviewStorage.Create(ctx, review); err != nil {
		s.logger.Error("failed to create review in storage", zap.Error(err))
		return nil, err
	}

	return s.getReviewResponse(ctx, review.ID)
}

func (s *Service) UpdateReview(ctx context.Context, id string, req dto.UpdateReviewRequest) (*dto.ReviewResponse, error) {
	review, err := s.reviewStorage.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	review.Rating = req.Rating
	review.Description = req.Description

	if err := s.reviewStorage.Update(ctx, review); err != nil {
		return nil, err
	}

	return s.getReviewResponse(ctx, id)
}

func (s *Service) DeleteReview(ctx context.Context, id string) error {
	return s.reviewStorage.Delete(ctx, id)
}

func (s *Service) GetReviewsByItemID(ctx context.Context, itemID string) ([]dto.ReviewResponse, error) {
	reviews, err := s.reviewStorage.GetByItemID(ctx, itemID)
	if err != nil {
		s.logger.Error("failed to get reviews by item id", zap.Error(err))
		return nil, err
	}

	res := make([]dto.ReviewResponse, len(reviews))
	for i, review := range reviews {
		res[i] = dto.ReviewResponse{
			ID:          review.ID,
			UserID:      review.UserID,
			ItemID:      review.ItemID,
			Rating:      review.Rating,
			Description: review.Description,
			CreatedAt:   review.CreatedAt,
			UpdatedAt:   review.UpdatedAt,
		}
	}

	return res, nil
}

func (s *Service) getReviewResponse(ctx context.Context, id string) (*dto.ReviewResponse, error) {
	review, err := s.reviewStorage.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get review by id", zap.Error(err))
		return nil, err
	}

	return &dto.ReviewResponse{
		ID:          review.ID,
		UserID:      review.UserID,
		ItemID:      review.ItemID,
		Rating:      review.Rating,
		Description: review.Description,
		CreatedAt:   review.CreatedAt,
		UpdatedAt:   review.UpdatedAt,
	}, nil
} 