package converter

import (
	"wowza/internal/dto"
	"wowza/internal/entity"
)

func ToItemResponse(item *entity.Item) *dto.ItemResponse {
	return &dto.ItemResponse{
		ID:          item.ID,
		BusinessID:  item.BusinessID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		ImageURL:    item.ImageURL,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func ToItemResponseList(items []entity.Item) []dto.ItemResponse {
	res := make([]dto.ItemResponse, len(items))
	for i := range items {
		res[i] = *ToItemResponse(&items[i])
	}
	return res
} 