package converter

import (
	"wowza/internal/dto"
	"wowza/internal/entity"
)

func ToBusinessResponse(business *entity.Business) *dto.BusinessResponse {
	return &dto.BusinessResponse{
		ID:          business.ID,
		UserID:      business.UserID,
		Name:        business.Name,
		Description: business.Description,
		WebsiteURL:  business.WebsiteURL,
		Location:    business.Location,
		Category:    ToCategoryResponse(business.Category),
	}
} 