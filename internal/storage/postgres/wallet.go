package postgres

import (
	"context"
	"errors"
	"wowza/internal/entity"

	"github.com/nordew/go-errx"
	"gorm.io/gorm"
)

type WalletStorage struct {
	db *gorm.DB
}

func NewWalletStorage(db *gorm.DB) *WalletStorage {
	return &WalletStorage{db: db}
}

func (s *WalletStorage) GetByUserID(ctx context.Context, userID string) (*entity.Wallet, error) {
	var wallet entity.Wallet

	if err := s.db.WithContext(ctx).Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewNotFound().WithDescription("wallet not found")
		}

		return nil, errx.NewInternal().WithDescription("failed to get wallet by user id")
	}

	return &wallet, nil
}

func (s *WalletStorage) Update(ctx context.Context, wallet *entity.Wallet) error {
	if err := s.db.WithContext(ctx).Save(wallet).Error; err != nil {
		return errx.NewInternal().WithDescription("failed to update wallet")
	}

	return nil
}
