package http

import (
	"context"
	"wowza/internal/dto"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) initPostRoutes(router fiber.Router) {
	posts := router.Group("/posts")
	{
		posts.Post("/", h.createPost)
	}
}

// @Summary Create Post
// @Tags posts
// @Description Creates a new post
// @Accept  multipart/form-data
// @Produce  json
// @Param   video          formData file   true  "Video file"
// @Param   user_id        formData string true  "User ID"
// @Param   description    formData string false "Post description"
// @Param   duration       formData number true  "Video duration"
// @Param   visibility     formData string true  "Post visibility" Enums(public, friends, private)
// @Param   hashtags       formData string false "Hashtags for the post (use multiple fields for multiple hashtags)"
// @Param   tags           formData string false "Tags for the post (use multiple fields for multiple tags)"
// @Param   allow_comments formData bool   true  "Allow comments"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/posts [post]
func (h *Handler) createPost(c fiber.Ctx) error {
	var req dto.CreatePostRequest
	if err := c.Bind().Form(&req); err != nil {
		return h.handleError(c, err)
	}

	fileHeader, err := c.FormFile("video")
	if err != nil {
		return h.handleError(c, err)
	}
	req.FileHeader = fileHeader

	ctx, cancel := context.WithTimeout(context.Background(), h.ctxTimeout)
	defer cancel()

	if err := h.services.Post.CreatePost(ctx, &req); err != nil {
		return h.handleError(c, err)
	}

	return h.handleSuccess(c, nil)
}
 