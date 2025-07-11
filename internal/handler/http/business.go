package http

import (
	"context"
	"wowza/internal/dto"

	"github.com/gofiber/fiber/v3"
)

type BusinessService interface {
	CreateBusiness(ctx context.Context, req dto.CreateBusinessRequest) (*dto.BusinessResponse, error)
	GetBusinessByID(ctx context.Context, id string) (*dto.BusinessResponse, error)
	UpdateBusiness(ctx context.Context, id string, req dto.UpdateBusinessRequest) (*dto.BusinessResponse, error)
	DeleteBusiness(ctx context.Context, id string) error
}

func (h *Handler) initBusinessRoutes(api fiber.Router) {
	businesses := api.Group("/businesses")
	{
		businesses.Post("/", h.createBusiness)
		businesses.Get("/:id", h.getBusinessByID)
		businesses.Put("/:id", h.updateBusiness)
		businesses.Delete("/:id", h.deleteBusiness)
	}
}

// @Summary Create Business
// @Description Create a new business page
// @Tags businesses
// @Accept json
// @Produce json
// @Param input body dto.CreateBusinessRequest true "Business Info"
// @Success 201 {object} dto.BusinessResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /businesses [post]
func (h *Handler) createBusiness(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	var req dto.CreateBusinessRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	business, err := h.services.Business.CreateBusiness(ctx, req)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(business)
}

// @Summary Get Business By ID
// @Description Get a business page by its ID
// @Tags businesses
// @Accept json
// @Produce json
// @Param id path string true "Business ID"
// @Success 200 {object} dto.BusinessResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /businesses/{id} [get]
func (h *Handler) getBusinessByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	id := c.Params("id")

	business, err := h.services.Business.GetBusinessByID(ctx, id)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, business)
}

// @Summary Update Business
// @Description Update a business page
// @Tags businesses
// @Accept json
// @Produce json
// @Param id path string true "Business ID"
// @Param input body dto.UpdateBusinessRequest true "Business Info"
// @Success 200 {object} dto.BusinessResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /businesses/{id} [put]
func (h *Handler) updateBusiness(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	id := c.Params("id")

	var req dto.UpdateBusinessRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	business, err := h.services.Business.UpdateBusiness(ctx, id, req)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, business)
}

// @Summary Delete Business
// @Description Delete a business page
// @Tags businesses
// @Accept json
// @Produce json
// @Param id path string true "Business ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /businesses/{id} [delete]
func (h *Handler) deleteBusiness(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	id := c.Params("id")

	if err := h.services.Business.DeleteBusiness(ctx, id); err != nil {
		return h.handleError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
} 