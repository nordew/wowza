package http

import (
	"context"
	"wowza/internal/dto"

	"github.com/gofiber/fiber/v3"
)

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error)
}

func (h *Handler) initCategoryRoutes(api fiber.Router) {
	categories := api.Group("/categories")
	{
		categories.Get("/", h.getAllCategories)
	}
}

// @Summary Get All Categories
// @Description Get a list of all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} dto.CategoryResponse
// @Failure 500 {object} map[string]string
// @Router /categories [get]
func (h *Handler) getAllCategories(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	categories, err := h.service.(CategoryService).GetAllCategories(ctx)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, categories)
} 