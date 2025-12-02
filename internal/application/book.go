package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type Book struct {
	bookRepo        domain.BookRepository
	bookService     domain.BookService
	categoryRepo    domain.CategoryRepository
	categoryService domain.CategoryService
	mediaRepo       domain.MediaRepository
	mediaService    domain.MediaService
}

func ProvideBook(
	bookRepo domain.BookRepository,
	bookService domain.BookService,
	categoryRepo domain.CategoryRepository,
	categoryService domain.CategoryService,
	mediaRepo domain.MediaRepository,
	mediaService domain.MediaService,
) *Book {
	return &Book{
		bookRepo:        bookRepo,
		bookService:     bookService,
		categoryRepo:    categoryRepo,
		categoryService: categoryService,
		mediaRepo:       mediaRepo,
		mediaService:    mediaService,
	}
}

var _ http.BookApplication = &Book{}

func (b *Book) List(ctx context.Context, param http.ListBookRequestDTO) (*http.PaginationResponseDTO[http.BookResponseDTO], error) {
	queryParams := param.QueryParams
	if queryParams == nil {
		queryParams = &http.ListBookRequestQueryParams{}
	}

	books, err := b.bookRepo.List(
		ctx,
		domain.BookRepositoryListParam{
			BookIDs:     queryParams.CategoryIDs,
			Search:      queryParams.Search,
			CategoryIDs: queryParams.CategoryIDs,
			Publisher:   queryParams.Publisher,
			Year:        queryParams.Year,
			MinPrice:    queryParams.MinPrice,
			MaxPrice:    queryParams.MaxPrice,
			SortRecent:  queryParams.SortRecent,
			SortPrice:   queryParams.SortPrice,
			Limit:       queryParams.Limit,
			Offset:      (queryParams.Page - 1) * queryParams.Limit,
		},
	)
	if err != nil {
		return nil, err
	}

	count, err := b.bookRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	bookDtos, err := b.enrichBooks(ctx, *books)
	if err != nil {
		return nil, err
	}

	pagination := newPaginationResponseDto(
		bookDtos,
		*count,
		queryParams.Page,
		queryParams.Limit,
	)
	return pagination, nil
}

func (b *Book) Get(ctx context.Context, param http.GetBookRequestDTO) (*http.BookResponseDTO, error) {
	book, err := b.bookRepo.Get(ctx, domain.BookRepositoryGetParam{
		BookID: param.PathParams.BookID,
	})
	if err != nil {
		return nil, err
	}

	books := []domain.Book{*book}
	enrichedBooks, err := b.enrichBooks(ctx, books)
	if err != nil {
		return nil, err
	}

	if len(enrichedBooks) == 0 {
		return nil, domain.ErrNotFound
	}

	return &enrichedBooks[0], nil
}

func (b *Book) Create(ctx context.Context, param http.CreateBookRequestDTO) (*http.BookResponseDTO, error) {
	// Validate categories exist
	categories, err := b.categoryRepo.List(
		ctx,
		domain.CategoryRepositoryListParam{
			CategoryIDs: param.Data.CategoryIDs,
		},
	)
	if err != nil {
		return nil, err
	}
	if len(*categories) != len(param.Data.CategoryIDs) {
		return nil, domain.ErrNotFound
	}

	mediaIDs := make([]uuid.UUID, 0, len(param.Data.Media))
	for _, m := range param.Data.Media {
		mediaIDs = append(mediaIDs, m.MediaID)
	}

	media, err := b.mediaRepo.List(
		ctx,
		domain.MediaRepositoryListParam{
			MediaIDs: mediaIDs,
		},
	)
	if err != nil {
		return nil, err
	}
	if len(*media) != len(param.Data.Media) {
		return nil, domain.ErrNotFound
	}

	// Convert request DTOs to domain BookMedia
	bookMedia := make([]domain.BookMedia, 0, len(param.Data.Media))
	for _, m := range param.Data.Media {
		bookMedia = append(bookMedia, domain.BookMedia{
			MediaID: m.MediaID,
			IsCover: m.IsCover,
		})
	}

	book, err := domain.NewBook(
		param.Data.Title,
		param.Data.Description,
		param.Data.Author,
		param.Data.Price,
		param.Data.PagesCount,
		param.Data.YearPublished,
		param.Data.Publisher,
		param.Data.Weight,
		param.Data.StockQuantity,
		param.Data.CategoryIDs,
		bookMedia,
	)
	if err != nil {
		return nil, err
	}

	if err := b.bookService.Validate(*book); err != nil {
		return nil, err
	}

	err = b.bookRepo.Save(ctx, domain.BookRepositorySaveParam{
		Book: *book,
	})
	if err != nil {
		return nil, err
	}

	books := []domain.Book{*book}
	enrichedBooks, err := b.enrichBooks(ctx, books)
	if err != nil {
		return nil, err
	}

	if len(enrichedBooks) == 0 {
		return nil, domain.ErrNotFound
	}

	return &enrichedBooks[0], nil
}

func (b *Book) Update(ctx context.Context, param http.UpdateBookRequestDTO) (*http.BookResponseDTO, error) {
	book, err := b.bookRepo.Get(ctx, domain.BookRepositoryGetParam{
		BookID: param.PathParams.BookID,
	})
	if err != nil {
		return nil, err
	}

	if len(param.Data.CategoryIDs) > 0 {
		categories, err := b.categoryRepo.List(
			ctx,
			domain.CategoryRepositoryListParam{
				CategoryIDs: param.Data.CategoryIDs,
			},
		)
		if err != nil {
			return nil, err
		}
		if len(*categories) != len(param.Data.CategoryIDs) {
			return nil, domain.ErrNotFound
		}
	}

	book.Update(
		param.Data.Title,
		param.Data.Description,
		param.Data.Author,
		param.Data.Price,
		param.Data.PagesCount,
		param.Data.YearPublished,
		param.Data.Publisher,
		param.Data.Weight,
		param.Data.StockQuantity,
		param.Data.CategoryIDs,
	)

	if err := b.bookService.Validate(*book); err != nil {
		return nil, err
	}

	err = b.bookRepo.Save(ctx, domain.BookRepositorySaveParam{
		Book: *book,
	})
	if err != nil {
		return nil, err
	}

	books := []domain.Book{*book}
	enrichedBooks, err := b.enrichBooks(ctx, books)
	if err != nil {
		return nil, err
	}

	if len(enrichedBooks) == 0 {
		return nil, domain.ErrNotFound
	}

	return &enrichedBooks[0], nil
}

func (b *Book) Delete(ctx context.Context, param http.DeleteBookRequestDTO) error {
	book, err := b.bookRepo.Get(ctx, domain.BookRepositoryGetParam{
		BookID: param.PathParams.BookID,
	})
	if err != nil {
		return err
	}

	book.Remove()

	if err := b.bookService.Validate(*book); err != nil {
		return err
	}

	err = b.bookRepo.Save(ctx, domain.BookRepositorySaveParam{
		Book: *book,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *Book) enrichBooks(ctx context.Context, books []domain.Book) ([]http.BookResponseDTO, error) {
	if len(books) == 0 {
		return []http.BookResponseDTO{}, nil
	}

	categoryIDsMap := make(map[string]struct{})
	mediaIDsMap := make(map[string]struct{})
	for _, book := range books {
		for _, catID := range book.CategoryIDs {
			categoryIDsMap[catID.String()] = struct{}{}
		}
		for _, bookMedia := range book.Media {
			mediaIDsMap[bookMedia.MediaID.String()] = struct{}{}
		}
	}

	var categories *[]domain.Category
	if len(categoryIDsMap) > 0 {
		var err error
		categories, err = b.categoryRepo.List(
			ctx,
			domain.CategoryRepositoryListParam{},
		)
		if err != nil {
			return nil, err
		}
	}

	var media *[]domain.Media
	if len(mediaIDsMap) > 0 {
		var err error
		media, err = b.mediaRepo.List(
			ctx,
			domain.MediaRepositoryListParam{},
		)
		if err != nil {
			return nil, err
		}
	}

	// Create maps for quick lookup
	categoryMap := make(map[string]*domain.Category)
	if categories != nil {
		for i := range *categories {
			categoryMap[(*categories)[i].ID.String()] = &(*categories)[i]
		}
	}

	mediaMap := make(map[string]*domain.Media)
	if media != nil {
		for i := range *media {
			mediaMap[(*media)[i].ID.String()] = &(*media)[i]
		}
	}

	bookDtos := make([]http.BookResponseDTO, 0, len(books))
	for _, book := range books {
		bookCategories := make([]*domain.Category, 0, len(book.CategoryIDs))
		for _, catID := range book.CategoryIDs {
			if cat, exists := categoryMap[catID.String()]; exists {
				bookCategories = append(bookCategories, cat)
			}
		}

		bookDto := http.ToBookResponseDTO(&book, bookCategories, mediaMap)
		if bookDto != nil {
			bookDtos = append(bookDtos, *bookDto)
		}
	}

	return bookDtos, nil
}
