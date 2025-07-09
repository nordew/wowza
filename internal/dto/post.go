package dto

import "mime/multipart"

type PostVisibility string

const (
	PostVisibilityPublic  PostVisibility = "public"
	PostVisibilityFriends PostVisibility = "friends"
	PostVisibilityPrivate PostVisibility = "private"
)

func (pv PostVisibility) IsValid() bool {
	switch pv {
	case PostVisibilityPublic, PostVisibilityFriends, PostVisibilityPrivate:
		return true
	default:
		return false
	}
}

type CreatePostRequest struct {
	UserID        string
	Description   string
	Duration      float64
	Visibility    PostVisibility
	Hashtags      []string
	Tags          []string
	AllowComments bool
	FileHeader    *multipart.FileHeader
}
