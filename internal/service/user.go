package service

import (
	"context"
	"wowza/internal/dto"
	"wowza/internal/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	initialWalletBalance = 0
	initialWalletCurrency = "USD"
)

func (s *Service) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*entity.User, error) {
	user, wallet, err := s.newUserWithWallet(req)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}

	if err := s.userStorage.CreateWithWallet(ctx, user, wallet); err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (s *Service) newUserWithWallet(req dto.CreateUserRequest) (*entity.User, *entity.Wallet, error) {
	userID := uuid.New().String()
	walletID := uuid.New().String()

	user, err := entity.NewUser(
		userID,
		req.ProfileName,
		req.FullName,
		req.Email,
		req.Phone,
		req.Password,
	)
	if err != nil {
		return nil, nil, err
	}

	wallet, err := entity.NewWallet(
		walletID,
		userID,
		initialWalletBalance,
		initialWalletCurrency,
	)
	if err != nil {
		return nil, nil, err
	}

	hashedPassword, err := s.passwordHasher.HashPassword(req.Password)
	if err != nil {
		return nil, nil, err
	}

	user.Password = hashedPassword

	return user, wallet, nil
}