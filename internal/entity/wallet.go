package entity

import "time"

type Wallet struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Balance   int64
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewWallet(
	id,
	userID string,
	balance int64,
	currency string,
) (*Wallet, error) {
	now := time.Now()

	return &Wallet{
		ID:        id,
		UserID:    userID,
		Balance:   balance,
		Currency:  currency,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}