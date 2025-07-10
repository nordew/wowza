package converter

import (
	"wowza/internal/dto"
	"wowza/internal/entity"
)

func ToReviewResponse(review *entity.Review) *dto.ReviewResponse {
	return &dto.ReviewResponse{
		ID:          review.ID,
		UserID:      review.UserID,
		ItemID:      review.ItemID,
		Rating:      review.Rating,
		Description: review.Description,
		CreatedAt:   review.CreatedAt,
		UpdatedAt:   review.UpdatedAt,
	}
}

func ToReviewResponseList(reviews []entity.Review) []dto.ReviewResponse {
	res := make([]dto.ReviewResponse, len(reviews))
	for i := range reviews {
		res[i] = *ToReviewResponse(&reviews[i])
	}
	return res
} 