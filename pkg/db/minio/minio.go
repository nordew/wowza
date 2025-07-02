package minio

import (
	"context"
	"log"

	"wowza/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func New(cfg config.Minio) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MinIO")

	err = createBucket(minioClient, cfg.BucketName)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func createBucket(client *minio.Client, bucketName string) error {
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
		log.Printf("Bucket %s created successfully", bucketName)
	}

	return nil
}
