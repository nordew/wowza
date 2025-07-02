package storage

import (
	"context"
	"net/url"
	"time"
	"wowza/internal/dto"

	"github.com/minio/minio-go/v7"
	"github.com/nordew/go-errx"
)

type Storage struct {
	client     *minio.Client
	bucketName string
}

func New(client *minio.Client, bucketName string) *Storage {
	return &Storage{
		client:     client,
		bucketName: bucketName,
	}
}

func (s *Storage) UploadFile(ctx context.Context, req dto.UploadFileRequest) (minio.UploadInfo, error) {
	info, err := s.client.PutObject(
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
		return info, errx.NewInternal().WithDescriptionAndCause("failed to upload file", err)
	}

	return info, nil
}

func (s *Storage) GetFilePresignedURL(
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

func (s *Storage) DeleteFile(ctx context.Context, objectName string) error {
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

func (s *Storage) GetFileInfo(
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
