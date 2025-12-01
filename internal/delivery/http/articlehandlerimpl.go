package http

import (
	"github.com/gin-gonic/gin"
)

type ArticleHandlerImpl struct{}

var _ ArticleHandler = &ArticleHandlerImpl{}

func ProvideArticleHandler() *ArticleHandlerImpl {
	return &ArticleHandlerImpl{}
}

// ListArticles godoc
//
//	@Summary		List all articles
//	@Description	Get all articles with optional search
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			queryParams	query	ListArticleRequestQueryParams	true	"Query parameters"
//	@Success		200			{object}	PaginationResponseDto[ArticleResponseDTO]
//	@Failure		500			{object}	Error
//	@Router			/articles [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ArticleHandlerImpl) List(ctx *gin.Context) {
}

// GetArticle godoc
//
//	@Summary		Get article by ID
//	@Description	Get article details by ID
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string	true	"Article ID"	format(uuid)
//	@Success		200			{object}	ArticleResponseDTO
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/articles/{id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ArticleHandlerImpl) Get(ctx *gin.Context) {
}

// CreateArticle godoc
//
//	@Summary		Create a new article
//	@Description	Create a new article
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			article	body		CreateArticleData	true	"Article request"
//	@Success		201		{object}	ArticleResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/articles [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ArticleHandlerImpl) Create(ctx *gin.Context) {
}

// UpdateArticle godoc
//
//	@Summary		Update article
//	@Description	Update article details
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Article ID"	format(uuid)
//	@Param			data	body		UpdateArticleData	true	"Update article request"
//	@Success		200		{object}	ArticleResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		404		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/articles/{id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ArticleHandlerImpl) Update(ctx *gin.Context) {
}

// DeleteArticle godoc
//
//	@Summary		Delete article
//	@Description	Delete article by ID
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Article ID"	format(uuid)
//	@Success		204
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/articles/{id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ArticleHandlerImpl) Delete(ctx *gin.Context) {
}
