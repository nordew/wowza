package converter

import (
	"wowza/internal/dto"
	"wowza/internal/entity"
)

func ToPostResponse(post entity.Post) dto.PostResponse {
	return dto.PostResponse{
		ID:            post.ID,
		UserID:        post.UserID,
		VideoURL:      post.VideoURL,
		Description:   post.Description,
		Duration:      post.Duration,
		Visibility:    post.Visibility,
		Hashtags:      post.Hashtags,
		Tags:          post.Tags,
		LikesCount:    post.LikesCount,
		CommentsCount: post.CommentsCount,
		ViewsCount:    post.ViewsCount,
		AllowComments: post.AllowComments,
		CreatedAt:     post.CreatedAt,
	}
}

func ToPostResponseList(posts []entity.Post) []dto.PostResponse {
	postDTOs := make([]dto.PostResponse, len(posts))
	for i, post := range posts {
		postDTOs[i] = ToPostResponse(post)
	}
	return postDTOs
} 