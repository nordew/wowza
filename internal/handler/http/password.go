package http

import (
	"context"
	"wowza/internal/dto"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) initPasswordRoutes(router fiber.Router) {
	password := router.Group("/password")
	{
		password.Post("/reset", h.resetPassword)
		password.Post("/reset/confirm", h.resetPasswordConfirm)
		password.Post("/reset/complete", h.resetPasswordConfirmComplete)
	}
}

// @Summary Reset password
// @Description Initiates the password reset process for a user.
// @Tags password
// @Accept json
// @Produce json
// @Param input body dto.ResetPasswordRequest true "User's email"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /api/v1/password/reset [post]
func (h *Handler) resetPassword(c fiber.Ctx) error {
	var req dto.ResetPasswordRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ctxTimeout)
	defer cancel()

	if err := h.service.ResetPassword(ctx, req); err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, nil)
}

// @Summary Confirm password reset
// @Description Confirms the password reset using a token sent to the user's email.
// @Tags password
// @Accept json
// @Produce json
// @Param input body dto.ResetPasswordConfirmRequest true "Reset token"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /api/v1/password/reset/confirm [post]
func (h *Handler) resetPasswordConfirm(c fiber.Ctx) error {
	var req dto.ResetPasswordConfirmRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ctxTimeout)
	defer cancel()

	if err := h.service.ResetPasswordConfirm(ctx, req); err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, nil)
}

// @Summary Complete password reset
// @Description Completes the password reset process by setting a new password.
// @Tags password
// @Accept json
// @Produce json
// @Param input body dto.ResetPasswordConfirmCompleteRequest true "New password and token"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /api/v1/password/reset/complete [post]
func (h *Handler) resetPasswordConfirmComplete(c fiber.Ctx) error {
	var req dto.ResetPasswordConfirmCompleteRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ctxTimeout)
	defer cancel()

	if err := h.service.ResetPasswordConfirmComplete(ctx, req); err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, nil)
}