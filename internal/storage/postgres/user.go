package postgres

import (
	"context"
	"errors"
	"wowza/internal/entity"

	"github.com/nordew/go-errx"
	"gorm.io/gorm"
)

type UserFilter struct {
	ID          string
	ProfileName string
	FullName    string
	Email       string
	Phone       string
}

type UserStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) Create(ctx context.Context, user *entity.User) error {
	if err := s.db.WithContext(ctx).Create(user).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to create user", err)
	}

	return nil
}

func (s *UserStorage) CreateWithWallet(
	ctx context.Context,
	user *entity.User,
	wallet *entity.Wallet,
) error {
	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return errx.NewInternal().WithDescriptionAndCause("failed to create user", err)
		}

		if err := tx.Create(wallet).Error; err != nil {
			return errx.NewInternal().WithDescriptionAndCause("failed to create wallet", err)
		}

		return nil
	}); err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to create user with wallet", err)
	}

	return nil
}

func (s *UserStorage) GetByFilter(ctx context.Context, filter UserFilter) (*entity.User, error) {
	var user entity.User

	if err := s.db.WithContext(ctx).Where(filter).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewNotFound().WithDescription("user not found")
		}

		return nil, errx.NewInternal().WithDescriptionAndCause("failed to get user by filter", err)
	}

	return &user, nil
}

func (s *UserStorage) Update(ctx context.Context, user *entity.User) error {
	if err := s.db.WithContext(ctx).Save(user).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to update user", err)
	}

	return nil
}

func (s *UserStorage) Delete(ctx context.Context, id string) error {
	if err := s.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to delete user", err)
	}

	return nil
}
