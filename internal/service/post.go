package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"wowza/internal/dto"
	"wowza/internal/entity"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
	"go.uber.org/zap"
)

func (s *Service) CreatePost(ctx context.Context, req *dto.CreatePostRequest) error {
	videoURL, err := s.uploadPostVideo(ctx, req.FileHeader)
	if err != nil {
		return err
	}

	post, err := entity.NewPost(uuid.NewString(), req, videoURL)
	if err != nil {
		s.logger.Error("failed to create new post entity", zap.Error(err))
		return err
	}

	if err := s.postStorage.Create(ctx, post); err != nil {
		s.logger.Error("failed to create post in storage", zap.Error(err))
		return err
	}

	return nil
}

func (s *Service) uploadPostVideo(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		s.logger.Error("failed to open file header", zap.Error(err))
		return "", errx.NewInternal().WithDescription("failed to process file")
	}
	defer file.Close()

	objectName := fmt.Sprintf("%s%s", uuid.NewString(), filepath.Ext(fileHeader.Filename))

	uploadReq := dto.UploadFileRequest{
		Name:        objectName,
		Reader:      file,
		Size:        fileHeader.Size,
		ContentType: fileHeader.Header.Get("Content-Type"),
	}

	if err := s.fileStorage.UploadFile(ctx, uploadReq); err != nil {
		s.logger.Error("failed to upload file to storage", zap.Error(err))
		return "", err
	}

	return s.fileStorage.GetFilePublicURL(objectName), nil
}
 