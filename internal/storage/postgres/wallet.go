package storage

import (
	"context"
	"wowza/internal/entity"

	"github.com/nordew/go-errx"
)

type GetWalletFilter struct {
	ID     string
	UserID string
}

func (s *Storage) CreateWallet(ctx context.Context, wallet *entity.Wallet) error {
	if err := s.db.WithContext(ctx).Create(wallet).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to create wallet", err)
	}

	return nil
}

func (s *Storage) GetWalletByFilter(ctx context.Context, filter GetWalletFilter) (*entity.Wallet, error) {
	var wallet entity.Wallet

	if err := s.db.WithContext(ctx).Where(filter).First(&wallet).Error; err != nil {
		return nil, errx.NewNotFound().WithDescription("wallet not found")
	}

	return &wallet, nil
}

func (s *Storage) UpdateWallet(ctx context.Context, wallet *entity.Wallet) error {
	if err := s.db.WithContext(ctx).Save(wallet).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to update wallet", err)
	}

	return nil
}

func (s *Storage) DeleteWallet(ctx context.Context, id string) error {
	if err := s.db.WithContext(ctx).Delete(&entity.Wallet{}, id).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to delete wallet", err)
	}

	return nil
}
