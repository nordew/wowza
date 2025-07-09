package postgres

import (
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

func (s *UserStorage) Create(user *entity.User) error {
	if err := s.db.Create(user).Error; err != nil {
		return errx.NewInternal().WithDescription("failed to create user")
	}

	return nil
}

func (s *UserStorage) CreateWithWallet(
	user *entity.User,
	wallet *entity.Wallet,
) error {
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		if err := tx.Create(wallet).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errx.NewInternal().WithDescription("failed to create user with wallet")
	}

	return nil
}

func (s *UserStorage) GetByFilter(filter UserFilter) (*entity.User, error) {
	var user entity.User

	if err := s.db.Where(filter).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewNotFound().WithDescription("user not found")
		}

		return nil, errx.NewInternal().WithDescription("failed to get user by filter")
	}

	return &user, nil
}

func (s *UserStorage) Update(user *entity.User) error {
	if err := s.db.Save(user).Error; err != nil {
		return errx.NewInternal().WithDescription("failed to update user")
	}

	return nil
}

func (s *UserStorage) Delete(id string) error {
	if err := s.db.Delete(&entity.User{}, id).Error; err != nil {
		return errx.NewInternal().WithDescription("failed to delete user")
	}

	return nil
}
