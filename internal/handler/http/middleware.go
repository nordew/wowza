package http

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func (h *Handler) loggingMiddleware(c fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	latency := time.Since(start)
	statusCode := c.Response().StatusCode()

	var fields []zap.Field
	fields = append(fields,
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.Int("status", statusCode),
		zap.Duration("latency", latency),
	)

	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	statusColor := colorForStatus(statusCode)
	methodColor := colorForMethod(c.Method())

	logMessage := fmt.Sprintf("%s %s %s | %3d | %13v | %s",
		statusColor,
		methodColor,
		c.Method(),
		statusCode,
		latency,
		c.Path(),
	)

	switch {
	case statusCode >= fiber.StatusInternalServerError:
		h.logger.Error(logMessage, fields...)
	case statusCode >= fiber.StatusBadRequest:
		h.logger.Warn(logMessage, fields...)
	default:
		h.logger.Info(logMessage, fields...)
	}

	return err
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "\x1b[32m[SUCCESS]\x1b[0m" // Green
	case code >= 300 && code < 400:
		return "\x1b[34m[REDIRECT]\x1b[0m" // Blue
	case code >= 400 && code < 500:
		return "\x1b[33m[CLIENT ERROR]\x1b[0m" // Yellow
	case code >= 500:
		return "\x1b[31m[SERVER ERROR]\x1b[0m" // Red
	default:
		return "\x1b[37m[UNKNOWN]\x1b[0m" // White
	}
}

func colorForMethod(method string) string {
	switch method {
	case fiber.MethodGet:
		return "\x1b[34m[GET]\x1b[0m" // Blue
	case fiber.MethodPost:
		return "\x1b[32m[POST]\x1b[0m" // Green
	case fiber.MethodPut:
		return "\x1b[33m[PUT]\x1b[0m" // Yellow
	case fiber.MethodDelete:
		return "\x1b[31m[DELETE]\x1b[0m" // Red
	case fiber.MethodPatch:
		return "\x1b[35m[PATCH]\x1b[0m" // Magenta
	default:
		return "\x1b[37m[" + method + "]\x1b[0m" // White
	}
}
