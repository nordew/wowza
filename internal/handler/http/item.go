package http

import (
	"context"
	"wowza/internal/dto"

	"github.com/gofiber/fiber/v3"
)

type ItemService interface {
	CreateItem(ctx context.Context, req dto.CreateItemRequest) (*dto.ItemResponse, error)
	GetItemByID(ctx context.Context, id string) (*dto.ItemResponse, error)
	UpdateItem(ctx context.Context, id string, req dto.UpdateItemRequest) (*dto.ItemResponse, error)
	DeleteItem(ctx context.Context, id string) error
	GetItemsByBusinessID(ctx context.Context, businessID string) ([]dto.ItemResponse, error)
}

func (h *Handler) initItemRoutes(api fiber.Router) {
	items := api.Group("/items")
	{
		items.Post("/", h.createItem)
		items.Get("/:id", h.getItemByID)
		items.Put("/:id", h.updateItem)
		items.Delete("/:id", h.deleteItem)
	}

	businessItems := api.Group("/businesses/:businessId/items")
	{
		businessItems.Get("/", h.getItemsByBusinessID)
	}
}

// Item handlers...
func (h *Handler) createItem(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	var req dto.CreateItemRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	item, err := h.service.(ItemService).CreateItem(ctx, req)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *Handler) getItemByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	id := c.Params("id")

	item, err := h.service.(ItemService).GetItemByID(ctx, id)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, item)
}

func (h *Handler) updateItem(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	id := c.Params("id")

	var req dto.UpdateItemRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	item, err := h.service.(ItemService).UpdateItem(ctx, id, req)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, item)
}

func (h *Handler) deleteItem(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	id := c.Params("id")

	if err := h.service.(ItemService).DeleteItem(ctx, id); err != nil {
		return h.handleError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) getItemsByBusinessID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	businessID := c.Params("businessId")

	items, err := h.service.(ItemService).GetItemsByBusinessID(ctx, businessID)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, items)
} 