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
		Duration:      int(post.Duration),
		Visibility:    string(post.Visibility),
		Hashtags:      post.Hashtags,
		Tags:          post.Tags,
		LikesCount:    int(post.LikesCount),
		CommentsCount: int(post.CommentsCount),
		ViewsCount:    int(post.ViewsCount),
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