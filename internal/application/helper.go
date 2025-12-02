package application

import (
	"math"

	"backend/internal/delivery/http"
)

func newPaginationResponseDto[T interface{}](
	data []T,
	totalItems, currentPage, itemsPerPage int,
) *http.PaginationResponseDTO[T] {
	if itemsPerPage <= 0 {
		itemsPerPage = 1
	}
	totalPages := int(math.Ceil(float64(totalItems) / float64(itemsPerPage)))
	return &http.PaginationResponseDTO[T]{
		Data: data,
		Meta: http.PaginationMetaResponseDTO{
			TotalPages:   int(totalPages),
			TotalItems:   totalItems,
			CurrentPage:  currentPage,
			ItemsPerPage: itemsPerPage,
			PageItems:    len(data),
		},
	}
}
