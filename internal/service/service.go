package service

import (
	"context"
	"time"
	"wowza/internal/entity"
	storage "wowza/internal/storage/postgres"
	"wowza/pkg/generator"

	"go.uber.org/zap"
)

type Storage interface {
	CreateUser(ctx context.Context, user *entity.User) error
	CreateUserWithWallet(ctx context.Context, user *entity.User, wallet *entity.Wallet) error
	GetUserByFilter(ctx context.Context, filter storage.UserFilter) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id string) error
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
	storage Storage
	logger  *zap.Logger
	passwordHasher PasswordHasher
	pasetoManager PasetoManager
	cache Cache
	generator Generator
}

func NewService(
	storage Storage,
	logger *zap.Logger,
	passwordHasher PasswordHasher,
	pasetoManager PasetoManager,
	cache Cache,
	generator Generator,
) *Service {
	return &Service{
		storage: storage,
		logger: logger,
		passwordHasher: passwordHasher,
		pasetoManager: pasetoManager,
		cache: cache,
		generator: generator,
	}
}