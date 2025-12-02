package http

type DeleteImageURLResponseDTO struct {
	URL string `json:"url" binding:"required,url"`
}

type UploadImageURLResponseDTO struct {
	URL string `json:"url" binding:"required,url"`
	Key string `json:"key" binding:"required"`
}

type PaginationMetaResponseDTO struct {
	TotalPages   int `json:"totalPages"   binding:"required"`
	TotalItems   int `json:"totalItems"   binding:"required"`
	CurrentPage  int `json:"currentPage"  binding:"required"`
	ItemsPerPage int `json:"itemsPerPage" binding:"required"`
	PageItems    int `json:"pageItems"`
}

type PaginationResponseDTO[T interface{}] struct {
	Data []T                       `json:"data" binding:"required"`
	Meta PaginationMetaResponseDTO `json:"meta" binding:"required"`
}
