package paseto

import (
	"errors"
	"time"
	"wowza/internal/entity"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

var (
	ErrTokenInvalid = errors.New("token is invalid")
	ErrTokenExpired = errors.New("token has expired")
)

type Manager struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewManager(symmetricKey []byte) *Manager {
	return &Manager{
		paseto:       paseto.NewV2(),
		symmetricKey: symmetricKey,
	}
}

func (m *Manager) CreateToken(user entity.User, duration time.Duration) (string, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	claims := &CustomClaims{
		User:      user,
		ID:        tokenID,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return m.paseto.Encrypt(m.symmetricKey, claims, nil)
}

func (m *Manager) VerifyToken(token string) (*entity.User, error) {
	claims := &CustomClaims{}

	err := m.paseto.Decrypt(token, m.symmetricKey, claims, nil)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	if err := claims.Valid(); err != nil {
		return nil, err
	}

	return &claims.User, nil
}
