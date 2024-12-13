package helper

import (
	"github.com/leftsidebrain/pzn-go-restful-api/model/domain"
	"github.com/leftsidebrain/pzn-go-restful-api/model/web"
)

func ToCagegoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCagegoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCagegoryResponse(category))
	}
	return categoryResponses
}