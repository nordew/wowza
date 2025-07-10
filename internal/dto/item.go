package dto

import "time"

type CreateItemRequest struct {
	BusinessID  string  `json:"businessId" binding:"required,uuid"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	ImageURL    string  `json:"imageUrl" binding:"omitempty,url"`
}

type UpdateItemRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price" binding:"omitempty,gt=0"`
	ImageURL    *string  `json:"imageUrl" binding:"omitempty,url"`
}

type ItemResponse struct {
	ID          string           `json:"id"`
	BusinessID  string           `json:"businessId"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	ImageURL    string           `json:"imageUrl"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
	Reviews     []ReviewResponse `json:"reviews,omitempty"`
} 