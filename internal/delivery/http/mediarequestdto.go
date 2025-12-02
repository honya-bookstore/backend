package http

type ListMediaRequestDTO struct {
	QueryParams *ListMediaRequestQueryParams
}

type ListMediaRequestQueryParams struct {
	PaginationRequestDTO
	Search string `json:"search"`
}

type GetMediaRequestDTO struct {
	PathParams *GetMediaRequestPathParams
}

type GetMediaRequestPathParams struct {
	MediaID string `json:"id" binding:"required"`
}

type CreateMediaRequestDTO struct {
	Data *CreateMediaRequestData `json:"data" binding:"required,dive"`
}

type CreateMediaRequestData struct {
	URL     string `json:"url"     binding:"required,url"`
	AltText string `json:"altText" binding:"omitempty,lte=200"`
	Order   int    `json:"order"   binding:"required,gte=0"`
}

type DeleteMediaRequestDTO struct {
	PathParams *DeleteMediaRequestPathParams
}

type DeleteMediaRequestPathParams struct {
	MediaID string `json:"id" binding:"required"`
}
