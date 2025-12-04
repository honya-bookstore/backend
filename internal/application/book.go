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

	// Convert to DTOs with enrichment
	bookDtos := make([]http.BookResponseDTO, 0, len(*books))
	if len(*books) > 0 {
		// Collect unique category and media IDs
		categoryIDsMap := make(map[string]struct{})
		mediaIDsMap := make(map[string]struct{})
		for _, book := range *books {
			for _, catID := range book.CategoryIDs {
				categoryIDsMap[catID.String()] = struct{}{}
			}
			for _, bookMedia := range book.Medium {
				mediaIDsMap[bookMedia.MediaID.String()] = struct{}{}
			}
		}

		// Fetch categories and media
		var categories *[]domain.Category
		if len(categoryIDsMap) > 0 {
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
			media, err = b.mediaRepo.List(
				ctx,
				domain.MediaRepositoryListParam{},
			)
			if err != nil {
				return nil, err
			}
		}

		// Create lookup maps
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

		// Convert to DTOs
		for _, book := range *books {
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

	// Fetch categories
	categoryIDsMap := make(map[string]struct{})
	for _, catID := range book.CategoryIDs {
		categoryIDsMap[catID.String()] = struct{}{}
	}

	var categories *[]domain.Category
	if len(categoryIDsMap) > 0 {
		categories, err = b.categoryRepo.List(
			ctx,
			domain.CategoryRepositoryListParam{},
		)
		if err != nil {
			return nil, err
		}
	}

	// Fetch media
	mediaIDsMap := make(map[string]struct{})
	for _, bookMedia := range book.Medium {
		mediaIDsMap[bookMedia.MediaID.String()] = struct{}{}
	}

	var media *[]domain.Media
	if len(mediaIDsMap) > 0 {
		media, err = b.mediaRepo.List(
			ctx,
			domain.MediaRepositoryListParam{},
		)
		if err != nil {
			return nil, err
		}
	}

	// Create lookup maps
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

	// Convert to DTO
	bookCategories := make([]*domain.Category, 0, len(book.CategoryIDs))
	for _, catID := range book.CategoryIDs {
		if cat, exists := categoryMap[catID.String()]; exists {
			bookCategories = append(bookCategories, cat)
		}
	}

	bookDto := http.ToBookResponseDTO(book, bookCategories, mediaMap)
	if bookDto == nil {
		return nil, domain.ErrNotFound
	}

	return bookDto, nil
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

	bookMedium := make([]domain.BookMedia, 0, len(param.Data.Media))
	for i, m := range param.Data.Media {
		bookMedia := domain.NewBookMedia(
			m.MediaID,
			m.IsCover,
			i+1,
		)
		bookMedium = append(bookMedium, *bookMedia)
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
		bookMedium,
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

	// Create lookup maps for DTO conversion
	categoryMap := make(map[string]*domain.Category)
	for i := range *categories {
		categoryMap[(*categories)[i].ID.String()] = &(*categories)[i]
	}

	mediaMap := make(map[string]*domain.Media)
	for i := range *media {
		mediaMap[(*media)[i].ID.String()] = &(*media)[i]
	}

	// Convert to DTO
	bookCategories := make([]*domain.Category, 0, len(book.CategoryIDs))
	for _, catID := range book.CategoryIDs {
		if cat, exists := categoryMap[catID.String()]; exists {
			bookCategories = append(bookCategories, cat)
		}
	}

	bookDto := http.ToBookResponseDTO(book, bookCategories, mediaMap)
	if bookDto == nil {
		return nil, domain.ErrNotFound
	}

	return bookDto, nil
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

	bookMedium := make([]domain.BookMedia, 0, len(param.Data.Media))
	for i, m := range param.Data.Media {
		bookMedia := domain.NewBookMedia(
			m.MediaID,
			m.IsCover,
			i+1,
		)
		bookMedium = append(bookMedium, *bookMedia)
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
		bookMedium,
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

	// Fetch categories for DTO
	categoryIDsMap := make(map[string]struct{})
	for _, catID := range book.CategoryIDs {
		categoryIDsMap[catID.String()] = struct{}{}
	}

	var categories *[]domain.Category
	if len(categoryIDsMap) > 0 {
		categories, err = b.categoryRepo.List(
			ctx,
			domain.CategoryRepositoryListParam{},
		)
		if err != nil {
			return nil, err
		}
	}

	// Fetch media for DTO
	mediaIDsMap := make(map[string]struct{})
	for _, bookMedia := range book.Medium {
		mediaIDsMap[bookMedia.MediaID.String()] = struct{}{}
	}

	var media *[]domain.Media
	if len(mediaIDsMap) > 0 {
		media, err = b.mediaRepo.List(
			ctx,
			domain.MediaRepositoryListParam{},
		)
		if err != nil {
			return nil, err
		}
	}

	// Create lookup maps
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

	// Convert to DTO
	bookCategories := make([]*domain.Category, 0, len(book.CategoryIDs))
	for _, catID := range book.CategoryIDs {
		if cat, exists := categoryMap[catID.String()]; exists {
			bookCategories = append(bookCategories, cat)
		}
	}

	bookDto := http.ToBookResponseDTO(book, bookCategories, mediaMap)
	if bookDto == nil {
		return nil, domain.ErrNotFound
	}

	return bookDto, nil
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
