package http

type PaginationRequestDTO struct {
	Page  int `json:"page"  binding:"gte=1,lte=50"`
	Limit int `json:"limit" binding:"gte=1,lte=100"`
}
