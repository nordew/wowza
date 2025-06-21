package service

import (
	"context"
	"wowza/internal/dto"
	"wowza/internal/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *Service) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*entity.User, error) {
	user, err := s.newUser(req)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}

	if err := s.storage.CreateUser(ctx, user); err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (s *Service) newUser(req dto.CreateUserRequest) (*entity.User, error) {
	id := uuid.New().String()

	user, err := entity.NewUser(
		id,
		req.ProfileName,
		req.FullName,
		req.Email,
		req.Phone,
		req.Password,
	)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := s.passwordHasher.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	return user, nil
}