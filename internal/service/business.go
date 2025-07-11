package service

import (
	"context"
	"wowza/internal/converter"
	"wowza/internal/dto"
	"wowza/internal/entity"
	"wowza/internal/storage"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type BusinessService struct {
	businessStorage storage.Business
	logger          *zap.Logger
}

func NewBusinessService(deps Dependencies) *BusinessService {
	return &BusinessService{
		businessStorage: deps.Storages.Business,
		logger:          deps.Logger,
	}
}

func (s *BusinessService) CreateBusiness(ctx context.Context, req dto.CreateBusinessRequest) (*dto.BusinessResponse, error) {
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		s.logger.Error("failed to get user id from context")
		return nil, &serviceErr{msg: "unauthorized", code: 401}
	}

	business, err := entity.NewBusiness(
		uuid.NewString(),
		userID,
		req.Name,
		req.Description,
		req.WebsiteURL,
		req.Location,
		req.CategoryID,
	)
	if err != nil {
		s.logger.Error("failed to create new business entity", zap.Error(err))
		return nil, err
	}

	if err := s.businessStorage.Create(ctx, business, nil); err != nil {
		s.logger.Error("failed to create business in storage", zap.Error(err))
		return nil, err
	}

	return s.GetBusinessByID(ctx, business.ID)
}

func (s *BusinessService) GetBusinessByID(ctx context.Context, id string) (*dto.BusinessResponse, error) {
	business, err := s.businessStorage.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get business by id", zap.Error(err))
		return nil, err
	}

	return converter.ToBusinessResponse(business), nil
}

func (s *BusinessService) UpdateBusiness(ctx context.Context, id string, req dto.UpdateBusinessRequest) (*dto.BusinessResponse, error) {
	business, err := s.businessStorage.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		business.Name = *req.Name
	}
	if req.Description != nil {
		business.Description = *req.Description
	}
	if req.WebsiteURL != nil {
		business.WebsiteURL = *req.WebsiteURL
	}
	if req.Location != nil {
		business.Location = *req.Location
	}
	if req.CategoryID != nil {
		business.CategoryID = *req.CategoryID
	}

	if err := s.businessStorage.Update(ctx, business, nil); err != nil {
		return nil, err
	}

	return s.GetBusinessByID(ctx, id)
}

func (s *BusinessService) DeleteBusiness(ctx context.Context, id string) error {
	return s.businessStorage.Delete(ctx, id)
}

type serviceErr struct {
	msg  string
	code int
}

func (e *serviceErr) Error() string {
	return e.msg
} 