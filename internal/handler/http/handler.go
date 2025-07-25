package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
	_ "wowza/docs" // swagger docs
	"wowza/internal/service"

	swagger "github.com/Flussen/swagger-fiber-v3"
	"github.com/gofiber/fiber/v3"
	"github.com/nordew/go-errx"
	"go.uber.org/zap"
)

var errxCodeToHTTPStatus = map[errx.Code]int{
	errx.BadRequest:    http.StatusBadRequest,
	errx.Unauthorized:  http.StatusUnauthorized,
	errx.Forbidden:     http.StatusForbidden,
	errx.NotFound:      http.StatusNotFound,
	errx.Conflict:      http.StatusConflict,
	errx.AlreadyExists: http.StatusConflict,
	errx.Validation:    http.StatusUnprocessableEntity,
	errx.Internal:      http.StatusInternalServerError,
	errx.Timeout:       http.StatusGatewayTimeout,
}

type Handler struct {
	logger     *zap.Logger
	services   *service.Services
	ctxTimeout time.Duration
}

func NewHandler(
	logger *zap.Logger,
	services *service.Services,
	ctxTimeout time.Duration,
) *Handler {
	return &Handler{
		logger:     logger,
		services:   services,
		ctxTimeout: ctxTimeout,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	router := fiber.New()

	router.Use(h.loggingMiddleware)

	// Swagger
	router.Get("/swagger/*", swagger.HandlerDefault)

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *fiber.App) {
	api := router.Group("/api/v1")
	{
		h.initAuthRoutes(api)
		h.initPasswordRoutes(api)
		h.initPostRoutes(api)
		h.initBusinessRoutes(api)
		h.initCategoryRoutes(api)
		h.initItemRoutes(api)
		h.initReviewRoutes(api)
		h.initFeedRoutes(api)
	}
}

func (h *Handler) handleError(c fiber.Ctx, err error) error {
	const internalErrMsg = "internal server error"

	h.logger.Error("error handled", zap.Error(err))

	var (
		syntaxErr    *json.SyntaxError
		unmarshalErr *json.UnmarshalTypeError
	)

	if errors.As(err, &syntaxErr) || errors.As(err, &unmarshalErr) || errors.Is(err, io.EOF) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	code := errx.GetCode(err)
	if status, ok := errxCodeToHTTPStatus[code]; ok {
		return c.Status(status).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": internalErrMsg})
}

func (h *Handler) handleSuccess(c fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(data)
}
