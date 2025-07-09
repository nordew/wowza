package minio

import (
	"context"
	"fmt"
	"net/url"
	"time"
	"wowza/internal/config"
	"wowza/internal/dto"

	"github.com/minio/minio-go/v7"
	"github.com/nordew/go-errx"
)

type FileStorage struct {
	client     *minio.Client
	bucketName string
	endpoint   string
	useSSL     bool
}

func NewFileStorage(client *minio.Client, cfg config.Minio) *FileStorage {
	return &FileStorage{
		client:     client,
		bucketName: cfg.BucketName,
		endpoint:   cfg.Endpoint,
		useSSL:     cfg.UseSSL,
	}
}

func (s *FileStorage) UploadFile(ctx context.Context, req dto.UploadFileRequest) error {
	_, err := s.client.PutObject(
		ctx,
		s.bucketName,
		req.Name,
		req.Reader,
		req.Size,
		minio.PutObjectOptions{
			ContentType: req.ContentType,
		},
	)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to upload file", err)
	}

	return nil
}

func (s *FileStorage) GetFilePresignedURL(
	ctx context.Context,
	objectName string,
	expiry time.Duration,
) (*url.URL, error) {
	presignedURL, err := s.client.PresignedGetObject(
		ctx,
		s.bucketName,
		objectName,
		expiry,
		nil,
	)
	if err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause("failed to generate presigned URL", err)
	}

	return presignedURL, nil
}

func (s *FileStorage) DeleteFile(ctx context.Context, objectName string) error {
	err := s.client.RemoveObject(
		ctx,
		s.bucketName,
		objectName,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to delete file", err)
	}

	return nil
}

func (s *FileStorage) GetFileInfo(
	ctx context.Context,
	objectName string,
) (minio.ObjectInfo, error) {
	info, err := s.client.StatObject(
		ctx,
		s.bucketName,
		objectName,
		minio.StatObjectOptions{},
	)
	if err != nil {
		return minio.ObjectInfo{}, errx.NewInternal().WithDescriptionAndCause("failed to get file info", err)
	}

	return info, nil
}

func (s *FileStorage) GetFilePublicURL(objectName string) string {
	scheme := "http"
	if s.useSSL {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s/%s/%s", scheme, s.endpoint, s.bucketName, objectName)
}
