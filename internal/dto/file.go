package dto

import "io"

type UploadFileRequest struct {
	Name        string
	Reader      io.Reader
	Size        int64
	ContentType string
}
