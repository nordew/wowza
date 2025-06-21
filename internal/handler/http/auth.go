package http

import (
	"context"
	"wowza/internal/dto"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) initAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	{
		signUp := auth.Group("/sign-up")
		{
			signUp.Post("/init", h.signUpInit)
			signUp.Post("/verify", h.signUpVerify)
			signUp.Post("/complete", h.signUpComplete)
		}

		auth.Post("/sign-in", h.signIn)
	}
}

// @Summary SignUp Init
// @Tags auth
// @Description Creates a new user
// @Accept  json
// @Produce  json
// @Param input body dto.SignUpInitRequest true "Sign Up Init"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/sign-up/init [post]
func (h *Handler) signUpInit(c fiber.Ctx) error {
	var req dto.SignUpInitRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ctxTimeout)
	defer cancel()

	if err := h.service.SignUpInit(ctx, req); err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, nil)
}

// @Summary SignUp Verify
// @Tags auth
// @Description Verifies a new user
// @Accept  json
// @Produce  json
// @Param input body dto.SignUpVerifyRequest true "Sign Up Verify"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/sign-up/verify [post]
func (h *Handler) signUpVerify(c fiber.Ctx) error {
	var req dto.SignUpVerifyRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ctxTimeout)
	defer cancel()

	if err := h.service.SignUpVerify(ctx, req); err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, nil)
}

// @Summary SignUp Complete
// @Tags auth
// @Description Completes a new user registration
// @Accept  json
// @Produce  json
// @Param input body dto.CreateUserRequest true "Sign Up Complete"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/sign-up/complete [post]
func (h *Handler) signUpComplete(c fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ctxTimeout)
	defer cancel()

	_, err := h.service.CreateUser(ctx, req)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, nil)
}

// @Summary SignIn
// @Tags auth
// @Description Logs a user in
// @Accept  json
// @Produce  json
// @Param input body dto.SignInRequest true "Sign In"
// @Success 200 {object} dto.SignInResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/sign-in [post]
func (h *Handler) signIn(c fiber.Ctx) error {
	var req dto.SignInRequest
	if err := c.Bind().Body(&req); err != nil {
		return h.handleError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ctxTimeout)
	defer cancel()

	res, err := h.service.SignIn(ctx, req)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, res)
}