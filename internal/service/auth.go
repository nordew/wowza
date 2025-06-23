package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wowza/internal/dto"
	storage "wowza/internal/storage/postgres"
	"wowza/pkg/generator"

	"github.com/nordew/go-errx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	verificationCodeLength     = 6
	verificationCodeCacheTTL   = 5 * time.Minute
	verificationStatusCacheTTL = 10 * time.Minute
)

func (s *Service) SignUpInit(ctx context.Context, req dto.SignUpInitRequest) error {
	_, err := s.storage.GetUserByFilter(ctx, storage.UserFilter{Phone: req.Phone})
	if err == nil {
		return errx.NewConflict().WithDescription("user with this phone number already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Error("failed to get user by filter", zap.Error(err))
		return err
	}

	code, err := s.generator.GenerateCode(verificationCodeLength, generator.NumbersOnly)
	if err != nil {
		s.logger.Error("failed to generate verification code", zap.Error(err))
		return errx.NewInternal().WithDescription("failed to generate verification code")
	}

	cacheKey := verificationCodeKey(req.Phone)
	go func() {
		if err := s.cache.Set(context.Background(), cacheKey, code, verificationCodeCacheTTL); err != nil {
			s.logger.Error("failed to set verification code in cache", zap.Error(err), zap.String("key", cacheKey))
		}
	}()

	s.logger.Info("verification code generated", zap.String("phone", req.Phone), zap.String("code", code))

	return nil
}

func (s *Service) SignIn(ctx context.Context, req dto.SignInRequest) (*dto.SignInResponse, error) {
	user, err := s.storage.GetUserByFilter(ctx, storage.UserFilter{Phone: req.Phone})
	if err != nil {
		if errx.GetCode(err) == errx.NotFound {
			return nil, errx.NewUnauthorized().WithDescription("invalid credentials")
		}

		s.logger.Error("failed to get user by filter", zap.Error(err))
		return nil, err
	}

	if user.Blocked {
		return nil, errx.NewForbidden().WithDescription("user is blocked")
	}

	if !s.passwordHasher.CheckPasswordHash(req.Password, user.Password) {
		return nil, errx.NewUnauthorized().WithDescription("invalid credentials")
	}

	accessToken, err := s.pasetoManager.CreateToken(*user, 15*time.Minute)
	if err != nil {
		s.logger.Error("failed to create access token", zap.Error(err))
		return nil, errx.NewInternal().WithDescription("failed to create access token")
	}

	return &dto.SignInResponse{
		AccessToken: accessToken,
	}, nil
}

func (s *Service) SignUpVerify(ctx context.Context, req dto.SignUpVerifyRequest) error {
	var cachedCode string
	cacheKey := verificationCodeKey(req.Phone)

	if err := s.cache.Get(ctx, cacheKey, &cachedCode); err != nil {
		if errx.GetCode(err) == errx.NotFound {
			return errx.NewBadRequest().WithDescription("invalid or expired verification code")
		}
		s.logger.Error("failed to get verification code from cache", zap.Error(err))
		return err
	}

	if cachedCode != req.Code {
		return errx.NewBadRequest().WithDescription("invalid verification code")
	}

	go func() {
		if err := s.cache.Delete(context.Background(), cacheKey); err != nil {
			s.logger.Warn("failed to delete verification code from cache", zap.Error(err), zap.String("key", cacheKey))
		}
	}()

	statusCacheKey := verificationStatusKey(req.Phone)
	go func() {
		if err := s.cache.Set(context.Background(), statusCacheKey, true, verificationStatusCacheTTL); err != nil {
			s.logger.Error("failed to set verification status in cache", zap.Error(err), zap.String("key", statusCacheKey))
		}
	}()

	return nil
}

func verificationCodeKey(phone string) string {
	return fmt.Sprintf("verification_code:%s", phone)
}

func verificationStatusKey(phone string) string {
	return fmt.Sprintf("verification_status:%s", phone)
}
