package http

import (
	"context"
	"strconv"
	"wowza/internal/dto"

	"github.com/gofiber/fiber/v3"
)

type FeedService interface {
	GetFeed(ctx context.Context, cursor string, limit int) (*dto.FeedResponse, error)
}

func (h *Handler) initFeedRoutes(api fiber.Router) {
	feed := api.Group("/feed")
	{
		feed.Get("/", h.getFeed)
	}
}

// @Summary Get Feed
// @Description Get a feed of posts
// @Tags feed
// @Accept json
// @Produce json
// @Param cursor query string false "Cursor for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {object} dto.FeedResponse
// @Failure 500 {object} map[string]string
// @Router /feed [get]
func (h *Handler) getFeed(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), h.ctxTimeout)
	defer cancel()

	cursor := c.Query("cursor")
	limit, _ := strconv.Atoi(c.Query("limit"))

	feed, err := h.service.GetFeed(ctx, cursor, limit)
	if err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, feed)
} 