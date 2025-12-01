package http

type DeleteImageURLResponseDto struct {
	URL string `json:"url" binding:"required,url"`
}

type UploadImageURLResponseDto struct {
	URL string `json:"url" binding:"required,url"`
	Key string `json:"key" binding:"required"`
}

type PaginationMetaResponseDto struct {
	TotalItems   int `json:"totalItems"   binding:"required"`
	CurrentPage  int `json:"currentPage"  binding:"required"`
	ItemsPerPage int `json:"itemsPerPage" binding:"required"`
	PageItems    int `json:"pageItems"`
}

type PaginationResponseDto[T interface{}] struct {
	Data []T                       `json:"data" binding:"required"`
	Meta PaginationMetaResponseDto `json:"meta" binding:"required"`
}
