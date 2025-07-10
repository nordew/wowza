package service

import (
	"context"
	"time"
	"wowza/internal/dto"
	"wowza/internal/entity"
	"wowza/internal/storage/postgres"
	"wowza/pkg/generator"

	"go.uber.org/zap"
)

type UserStorage interface {
	Create(ctx context.Context, user *entity.User) error
	CreateWithWallet(ctx context.Context, user *entity.User, wallet *entity.Wallet) error
	GetByFilter(ctx context.Context, filter postgres.UserFilter) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}

type PostStorage interface {
	Create(ctx context.Context, post *entity.Post) error
	GetForFeed(ctx context.Context, cursor time.Time, limit int) ([]entity.Post, error)
}

type WalletStorage interface {
	GetByUserID(ctx context.Context, userID string) (*entity.Wallet, error)
	Update(ctx context.Context, wallet *entity.Wallet) error
}

type BusinessStorage interface {
	Create(ctx context.Context, business *entity.Business, categoryIDs []string) error
	GetByID(ctx context.Context, id string) (*entity.Business, error)
	Update(ctx context.Context, business *entity.Business, categoryIDs []string) error
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string) ([]entity.Business, error)
}

type CategoryStorage interface {
	GetAll(ctx context.Context) ([]entity.Category, error)
}

type ItemStorage interface {
	Create(ctx context.Context, item *entity.Item) error
	GetByID(ctx context.Context, id string) (*entity.Item, error)
	Update(ctx context.Context, item *entity.Item) error
	Delete(ctx context.Context, id string) error
	GetByBusinessID(ctx context.Context, businessID string) ([]entity.Item, error)
}

type ReviewStorage interface {
	Create(ctx context.Context, review *entity.Review) error
	GetByID(ctx context.Context, id string) (*entity.Review, error)
	Update(ctx context.Context, review *entity.Review) error
	Delete(ctx context.Context, id string) error
	GetByItemID(ctx context.Context, itemID string) ([]entity.Review, error)
}

type FileStorage interface {
	UploadFile(ctx context.Context, req dto.UploadFileRequest) error
	GetFilePublicURL(objectName string) string
}

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

type Service struct {
	userStorage     UserStorage
	postStorage     PostStorage
	walletStorage   WalletStorage
	businessStorage BusinessStorage
	categoryStorage CategoryStorage
	itemStorage     ItemStorage
	reviewStorage   ReviewStorage
	logger          *zap.Logger
	passwordHasher  PasswordHasher
	pasetoManager   PasetoManager
	cache           Cache
	generator       Generator
	fileStorage     FileStorage
}

func NewService(
	userStorage UserStorage,
	postStorage PostStorage,
	walletStorage WalletStorage,
	businessStorage BusinessStorage,
	categoryStorage CategoryStorage,
	itemStorage ItemStorage,
	reviewStorage ReviewStorage,
	logger *zap.Logger,
	passwordHasher PasswordHasher,
	pasetoManager PasetoManager,
	cache Cache,
	generator Generator,
	fileStorage FileStorage,
) *Service {
	return &Service{
		userStorage:     userStorage,
		postStorage:     postStorage,
		walletStorage:   walletStorage,
		businessStorage: businessStorage,
		categoryStorage: categoryStorage,
		itemStorage:     itemStorage,
		reviewStorage:   reviewStorage,
		logger:          logger,
		passwordHasher:  passwordHasher,
		pasetoManager:   pasetoManager,
		cache:           cache,
		generator:       generator,
		fileStorage:     fileStorage,
	}
}
