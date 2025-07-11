package storage

import (
	"context"
	"time"
	"wowza/internal/dto"
	"wowza/internal/entity"
	"wowza/internal/storage/minio"
	"wowza/internal/storage/postgres"

	"wowza/internal/config"

	minioclient "github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

// Storage interfaces that define the contract for each storage type
type User interface {
	Create(ctx context.Context, user *entity.User) error
	CreateWithWallet(ctx context.Context, user *entity.User, wallet *entity.Wallet) error
	GetByFilter(ctx context.Context, filter postgres.UserFilter) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}

type Post interface {
	Create(ctx context.Context, post *entity.Post) error
	GetForFeed(ctx context.Context, cursor time.Time, limit int) ([]entity.Post, error)
}

type Wallet interface {
	GetByUserID(ctx context.Context, userID string) (*entity.Wallet, error)
	Update(ctx context.Context, wallet *entity.Wallet) error
}

type Business interface {
	Create(ctx context.Context, business *entity.Business, categoryIDs []string) error
	GetByID(ctx context.Context, id string) (*entity.Business, error)
	Update(ctx context.Context, business *entity.Business, categoryIDs []string) error
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string) ([]entity.Business, error)
}

type Category interface {
	GetAll(ctx context.Context) ([]entity.Category, error)
}

type Item interface {
	Create(ctx context.Context, item *entity.Item) error
	GetByID(ctx context.Context, id string) (*entity.Item, error)
	Update(ctx context.Context, item *entity.Item) error
	Delete(ctx context.Context, id string) error
	GetByBusinessID(ctx context.Context, businessID string) ([]entity.Item, error)
}

type Review interface {
	Create(ctx context.Context, review *entity.Review) error
	GetByID(ctx context.Context, id string) (*entity.Review, error)
	Update(ctx context.Context, review *entity.Review) error
	Delete(ctx context.Context, id string) error
	GetByItemID(ctx context.Context, itemID string) ([]entity.Review, error)
}

type File interface {
	UploadFile(ctx context.Context, req dto.UploadFileRequest) error
	GetFilePublicURL(objectName string) string
}

// Storages struct that embeds all storage interfaces
type Storages struct {
	User     User
	Post     Post
	Wallet   Wallet
	Business Business
	Category Category
	Item     Item
	Review   Review
	File     File
}

// Dependencies struct for constructing storages
type Dependencies struct {
	DB          *gorm.DB
	MinioClient *minioclient.Client
	MinioConfig config.Minio
}

func NewStorages(deps Dependencies) *Storages {
	// Create postgres storages
	postgresStorages := postgres.NewStorages(deps.DB)
	
	// Create minio storage
	fileStorage := minio.NewFileStorage(deps.MinioClient, deps.MinioConfig)

	return &Storages{
		User:     postgresStorages.User,
		Post:     postgresStorages.Post,
		Wallet:   postgresStorages.Wallet,
		Business: postgresStorages.Business,
		Category: postgresStorages.Category,
		Item:     postgresStorages.Item,
		Review:   postgresStorages.Review,
		File:     fileStorage,
	}
} 