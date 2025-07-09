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
	Create(user *entity.User) error
	CreateWithWallet(user *entity.User, wallet *entity.Wallet) error
	GetByFilter(filter postgres.UserFilter) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id string) error
}

type PostStorage interface {
	Create(post *entity.Post) error
}

type WalletStorage interface {
	GetByUserID(userID string) (*entity.Wallet, error)
	Update(wallet *entity.Wallet) error
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
	userStorage    UserStorage
	postStorage    PostStorage
	walletStorage  WalletStorage
	logger         *zap.Logger
	passwordHasher PasswordHasher
	pasetoManager  PasetoManager
	cache          Cache
	generator      Generator
	fileStorage    FileStorage
}

func NewService(
	userStorage UserStorage,
	postStorage PostStorage,
	walletStorage WalletStorage,
	logger *zap.Logger,
	passwordHasher PasswordHasher,
	pasetoManager PasetoManager,
	cache Cache,
	generator Generator,
	fileStorage FileStorage,
) *Service {
	return &Service{
		userStorage:    userStorage,
		postStorage:    postStorage,
		walletStorage:  walletStorage,
		logger:         logger,
		passwordHasher: passwordHasher,
		pasetoManager:  pasetoManager,
		cache:          cache,
		generator:      generator,
		fileStorage:    fileStorage,
	}
}
