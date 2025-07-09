package postgres

import (
	"wowza/internal/entity"

	"gorm.io/gorm"
)

type User interface {
	Create(user *entity.User) error
	CreateWithWallet(user *entity.User, wallet *entity.Wallet) error
	GetByFilter(filter UserFilter) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id string) error
}

type Post interface {
	Create(post *entity.Post) error
}

type Wallet interface {
	GetByUserID(userID string) (*entity.Wallet, error)
	Update(wallet *entity.Wallet) error
}

type Storages struct {
	User   User
	Post   Post
	Wallet Wallet
}

func NewStorages(db *gorm.DB) *Storages {
	return &Storages{
		User:   NewUserStorage(db),
		Post:   NewPostStorage(db),
		Wallet: NewWalletStorage(db),
	}
} 