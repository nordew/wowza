package service

import (
	"context"
	"fmt"
	"time"
	"wowza/internal/dto"
	"wowza/internal/storage"
	storage_postgres "wowza/internal/storage/postgres"
	"wowza/pkg/generator"

	"github.com/nordew/go-errx"
	"go.uber.org/zap"
)

const (
	passwordResetCodeLength     = 6
	passwordResetCodeCacheTTL   = 5 * time.Minute
)

type PasswordService struct {
	userStorage    storage.User
	logger         *zap.Logger
	passwordHasher PasswordHasher
	cache          Cache
	generator      Generator
}

func NewPasswordService(deps Dependencies) *PasswordService {
	return &PasswordService{
		userStorage:    deps.Storages.User,
		logger:         deps.Logger,
		passwordHasher: deps.PasswordHasher,
		cache:          deps.Cache,
		generator:      deps.Generator,
	}
}

func (s *PasswordService) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {
	_, err := s.userStorage.GetByFilter(ctx, storage_postgres.UserFilter{
		Email: req.Email,
	})
	if err != nil {
		return err
	}

	code, err := s.generator.GenerateCode(passwordResetCodeLength, generator.NumbersOnly)
	if err != nil {
		s.logger.Error("failed to generate password reset code", zap.Error(err))
		return errx.NewInternal().WithDescription("failed to generate password reset code")
	}

	cacheKey := passwordResetCodeKey(req.Email)
	go func() {
		if err := s.cache.Set(context.Background(), cacheKey, code, passwordResetCodeCacheTTL); err != nil {
			s.logger.Error("failed to set password reset code in cache", zap.Error(err), zap.String("key", cacheKey))
		}
	}()

	s.logger.Info("password reset code generated", zap.String("email", req.Email), zap.String("code", code))

	return nil
}

func (s *PasswordService) ResetPasswordConfirm(ctx context.Context, req dto.ResetPasswordConfirmRequest) error {
	var cachedCode string
	cacheKey := passwordResetCodeKey(req.Email)

	if err := s.cache.Get(ctx, cacheKey, &cachedCode); err != nil {
		return err
	}

	if cachedCode != req.Code {
		return errx.NewBadRequest().WithDescription("invalid code")
	}

	return nil
}

func (s *PasswordService) ResetPasswordConfirmComplete(ctx context.Context, req dto.ResetPasswordConfirmCompleteRequest) error {
	user, err := s.userStorage.GetByFilter(ctx, storage_postgres.UserFilter{
		Email: req.Email,
	})
	if err != nil {
		return err
	}

	hashedPassword, err := s.passwordHasher.HashPassword(req.Password)
	if err != nil {
		s.logger.Error("failed to hash password", zap.Error(err))
		return errx.NewInternal().WithDescription("failed to hash password")
	}

	user.Password = hashedPassword
	if err := s.userStorage.Update(ctx, user); err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return errx.NewInternal().WithDescription("failed to update user")
	}

	return nil
}

func passwordResetCodeKey(email string) string {
	return fmt.Sprintf("password_reset_code:%s", email)
}