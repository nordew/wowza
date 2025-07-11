package service

import (
	"context"
	"time"
	"wowza/internal/dto"
	"wowza/internal/entity"
	"wowza/internal/storage"
	"wowza/pkg/generator"

	"go.uber.org/zap"
)

// External dependencies
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type PasetoManager interface {
	CreateToken(user entity.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*entity.User, error)
}

type Cache interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string, dest any) error
	Delete(ctx context.Context, key string) error
}

type Generator interface {
	GenerateCode(size int, charType generator.CharType) (string, error)
}

// Service interfaces
type Auth interface {
	SignUpInit(ctx context.Context, req dto.SignUpInitRequest) error
	SignUpVerify(ctx context.Context, req dto.SignUpVerifyRequest) error
	SignIn(ctx context.Context, req dto.SignInRequest) (*dto.SignInResponse, error)
}

type User interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*entity.User, error)
}

type Password interface {
	ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error
	ResetPasswordConfirm(ctx context.Context, req dto.ResetPasswordConfirmRequest) error
	ResetPasswordConfirmComplete(ctx context.Context, req dto.ResetPasswordConfirmCompleteRequest) error
}

type Post interface {
	CreatePost(ctx context.Context, req *dto.CreatePostRequest) error
}

type Business interface {
	CreateBusiness(ctx context.Context, req dto.CreateBusinessRequest) (*dto.BusinessResponse, error)
	GetBusinessByID(ctx context.Context, id string) (*dto.BusinessResponse, error)
	UpdateBusiness(ctx context.Context, id string, req dto.UpdateBusinessRequest) (*dto.BusinessResponse, error)
	DeleteBusiness(ctx context.Context, id string) error
}

type Category interface {
	GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error)
}

type Item interface {
	CreateItem(ctx context.Context, req dto.CreateItemRequest) (*dto.ItemResponse, error)
	GetItemByID(ctx context.Context, id string) (*dto.ItemResponse, error)
	UpdateItem(ctx context.Context, id string, req dto.UpdateItemRequest) (*dto.ItemResponse, error)
	DeleteItem(ctx context.Context, id string) error
	GetItemsByBusinessID(ctx context.Context, businessID string) ([]dto.ItemResponse, error)
}

type Review interface {
	CreateReview(ctx context.Context, req dto.CreateReviewRequest) (*dto.ReviewResponse, error)
	UpdateReview(ctx context.Context, id string, req dto.UpdateReviewRequest) (*dto.ReviewResponse, error)
	DeleteReview(ctx context.Context, id string) error
	GetReviewsByItemID(ctx context.Context, itemID string) ([]dto.ReviewResponse, error)
}

type Feed interface {
	GetFeed(ctx context.Context, cursor string, limit int) (*dto.FeedResponse, error)
}

// Services struct that embeds all service interfaces
type Services struct {
	Auth     Auth
	User     User
	Password Password
	Post     Post
	Business Business
	Category Category
	Item     Item
	Review   Review
	Feed     Feed
}

// Dependencies struct for constructing services
type Dependencies struct {
	Storages       *storage.Storages
	Logger         *zap.Logger
	PasswordHasher PasswordHasher
	PasetoManager  PasetoManager
	Cache          Cache
	Generator      Generator
}

func NewServices(deps Dependencies) *Services {
	return &Services{
		Auth:     NewAuthService(deps),
		User:     NewUserService(deps),
		Password: NewPasswordService(deps),
		Post:     NewPostService(deps),
		Business: NewBusinessService(deps),
		Category: NewCategoryService(deps),
		Item:     NewItemService(deps),
		Review:   NewReviewService(deps),
		Feed:     NewFeedService(deps),
	}
}
