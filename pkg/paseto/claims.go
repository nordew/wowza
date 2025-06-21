package paseto

import (
	"time"
	"wowza/internal/entity"

	"github.com/google/uuid"
)

type CustomClaims struct {
	User    entity.User `json:"user"`
	ID      uuid.UUID   `json:"jti"`
	IssuedAt  time.Time   `json:"iat"`
	ExpiresAt time.Time   `json:"exp"`
}

func (c *CustomClaims) Valid() error {
	if time.Now().After(c.ExpiresAt) {
		return ErrTokenExpired
	}
		
	return nil
}
