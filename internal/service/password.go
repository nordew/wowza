package service

import (
	"context"
	"fmt"
	"time"
	"wowza/internal/dto"
	storage "wowza/internal/storage/postgres"
	"wowza/pkg/generator"

	"github.com/nordew/go-errx"
	"go.uber.org/zap"
)

const (
	passwordResetCodeLength     = 6
	passwordResetCodeCacheTTL   = 5 * time.Minute
)

func (s *Service) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {
	_, err := s.userStorage.GetByFilter(storage.UserFilter{
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

func (s *Service) ResetPasswordConfirm(ctx context.Context, req dto.ResetPasswordConfirmRequest) error {
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

func (s *Service) ResetPasswordConfirmComplete(ctx context.Context, req dto.ResetPasswordConfirmCompleteRequest) error {
	user, err := s.userStorage.GetByFilter(storage.UserFilter{
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
	if err := s.userStorage.Update(user); err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return errx.NewInternal().WithDescription("failed to update user")
	}

	return nil
}


func passwordResetCodeKey(email string) string {
	return fmt.Sprintf("password_reset_code:%s", email)
}