package dto

import "time"

type CreateReviewRequest struct {
	ItemID      string `json:"itemId" binding:"required,uuid"`
	Rating      int    `json:"rating" binding:"required,min=1,max=5"`
	Description string `json:"description"`
}

type UpdateReviewRequest struct {
	Rating      int    `json:"rating" binding:"omitempty,min=1,max=5"`
	Description string `json:"description"`
}

type ReviewResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	ItemID      string    `json:"itemId"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
} 