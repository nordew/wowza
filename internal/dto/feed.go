package dto

import "time"

type GetFeedRequest struct {
	CursorStr string
	Limit     int
}

type GetFeedResponse struct {
	Posts      []PostResponse
	NextCursor string
}

type PostResponse struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	VideoURL      string    `json:"video_url"`
	Description   string    `json:"description"`
	Duration      int       `json:"duration"`
	Visibility    string    `json:"visibility"`
	Hashtags      []string  `json:"hashtags"`
	Tags          []string  `json:"tags"`
	LikesCount    int       `json:"likes_count"`
	CommentsCount int       `json:"comments_count"`
	ViewsCount    int       `json:"views_count"`
	AllowComments bool      `json:"allow_comments"`
	CreatedAt     time.Time `json:"created_at"`
}

type FeedResponse struct {
	Posts      []PostResponse `json:"posts"`
	NextCursor string         `json:"nextCursor,omitempty"`
} 