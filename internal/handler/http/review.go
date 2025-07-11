package http

import (
	"context"
	"wowza/internal/dto"

	"github.com/gofiber/fiber/v3"
)

type ReviewService interface {
	CreateReview(ctx context.Context, req dto.CreateReviewRequest) (*dto.ReviewResponse, error)
	UpdateReview(ctx context.Context, id string, req dto.UpdateReviewRequest) (*dto.ReviewResponse, error)
	DeleteReview(ctx context.Context, id string) error
	GetReviewsByItemID(ctx context.Context, itemID string) ([]dto.ReviewResponse, error)
}

func (h *Handler) initReviewRoutes(api fiber.Router) {
	reviews := api.Group("/reviews")
	{
		reviews.Post("/", h.createReview)
		reviews.Put("/:id", h.updateReview)
		reviews.Delete("/:id", h.deleteReview)
	}

	itemReviews := api.Group("/items/:itemId/reviews")
	{
		itemReviews.Get("/", h.getReviewsByItemID)
	}
}

func (h *Handler) createReview(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	var req dto.CreateReviewRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	review, err := h.services.Review.CreateReview(ctx, req)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(review)
}

func (h *Handler) updateReview(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	id := c.Params("id")

	var req dto.UpdateReviewRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	review, err := h.services.Review.UpdateReview(ctx, id, req)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, review)
}

func (h *Handler) deleteReview(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	id := c.Params("id")

	if err := h.services.Review.DeleteReview(ctx, id); err != nil {
		return h.handleError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) getReviewsByItemID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	itemID := c.Params("itemId")

	reviews, err := h.services.Review.GetReviewsByItemID(ctx, itemID)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, reviews)
} 