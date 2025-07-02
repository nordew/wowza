package dto

import "io"

type UploadFileRequest struct {
	Name        string
	Reader      io.Reader
	File        []byte
	ContentType string
	Size        int64
}
