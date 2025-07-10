package converter

import (
	"wowza/internal/dto"
	"wowza/internal/entity"
)

func ToCategoryResponse(category entity.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponseList(categories []entity.Category) []dto.CategoryResponse {
	res := make([]dto.CategoryResponse, len(categories))
	for i, category := range categories {
		res[i] = ToCategoryResponse(category)
	}
	return res
} 