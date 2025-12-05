package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Book struct {
	queries *sqlc.Queries
	conn    *pgxpool.Pool
}

var _ domain.BookRepository = (*Book)(nil)

func ProvideBook(
	queries *sqlc.Queries,
	conn *pgxpool.Pool,
) *Book {
	return &Book{
		queries: queries,
		conn:    conn,
	}
}

func (r *Book) List(
	ctx context.Context,
	params domain.BookRepositoryListParam,
) (*[]domain.Book, error) {
	bookEntities, err := r.queries.ListBooks(ctx, sqlc.ListBooksParams{
		IDs:         params.BookIDs,
		Search:      params.Search,
		CategoryIDs: params.CategoryIDs,
		Publisher:   params.Publisher,
		Year:        int32(params.Year),
		MinPrice:    int64ToNumeric(params.MinPrice),
		MaxPrice:    int64ToNumeric(params.MaxPrice),
		SortRecent:  params.SortRecent,
		SortPrice:   params.SortPrice,
		Limit:       int32(params.Limit),
		Offset:      int32(params.Offset),
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	books := make([]domain.Book, 0, len(bookEntities))
	for _, bookEntity := range bookEntities {
		books = append(books, *toBookDomain(&bookEntity))
	}
	err = r.fillBooksCategories(ctx, &books)
	if err != nil {
		return nil, toDomainError(err)
	}
	err = r.fillBooksMedium(ctx, &books)
	if err != nil {
		return nil, toDomainError(err)
	}
	return &books, nil
}

func (r *Book) Count(
	ctx context.Context,
) (*int, error) {
	count, err := r.queries.CountBooks(ctx, sqlc.CountBooksParams{})
	if err != nil {
		return nil, toDomainError(err)
	}
	countInt := int(count)
	return &countInt, nil
}

func (r *Book) Get(
	ctx context.Context,
	params domain.BookRepositoryGetParam,
) (*domain.Book, error) {
	bookEntity, err := r.queries.GetBook(ctx, sqlc.GetBookParams{
		ID: params.BookID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	book := toBookDomain(&bookEntity)
	books := []domain.Book{*book}
	err = r.fillBooksCategories(ctx, &books)
	if err != nil {
		return nil, toDomainError(err)
	}
	err = r.fillBooksMedium(ctx, &books)
	if err != nil {
		return nil, toDomainError(err)
	}
	return &books[0], nil
}

func (r *Book) Save(
	ctx context.Context,
	params domain.BookRepositorySaveParam,
) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return toDomainError(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	qtx := r.queries.WithTx(tx)
	err = qtx.UpsertBook(ctx, sqlc.UpsertBookParams{
		ID:            params.Book.ID,
		Title:         params.Book.Title,
		Description:   &params.Book.Description,
		Author:        params.Book.Author,
		Price:         int64ToNumeric(params.Book.Price),
		PagesCount:    int32(params.Book.PagesCount),
		Year:          int32(params.Book.YearPublished),
		Publisher:     params.Book.Publisher,
		Weight:        float64ToNumeric(params.Book.Weight),
		StockQuantity: int32(params.Book.StockQuantity),
		PurchaseCount: int32(params.Book.PurchaseCount),
		Rating:        float32(params.Book.Rating),
		CreatedAt: pgtype.Timestamptz{
			Time:  params.Book.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  params.Book.UpdatedAt,
			Valid: true,
		},
		DeletedAt: pgtype.Timestamptz{
			Time:  params.Book.DeletedAt,
			Valid: true,
		},
	})
	if err != nil {
		return toDomainError(err)
	}
	err = r.mergeCategories(ctx, qtx, &params.Book)
	if err != nil {
		return toDomainError(err)
	}
	err = r.mergeMedium(ctx, qtx, &params.Book)
	if err != nil {
		return toDomainError(err)
	}
	if err := tx.Commit(ctx); err != nil {
		return toDomainError(err)
	}
	return nil
}

func toBookDomain(
	bookEntity *sqlc.Book,
) *domain.Book {
	return &domain.Book{
		ID:            bookEntity.ID,
		Title:         bookEntity.Title,
		Description:   ptr.Deref(bookEntity.Description, ""),
		Author:        bookEntity.Author,
		Price:         numericToInt64(bookEntity.Price),
		PagesCount:    int(bookEntity.PagesCount),
		YearPublished: int(bookEntity.Year),
		Publisher:     bookEntity.Publisher,
		Weight:        numericToFloat64(bookEntity.Weight),
		StockQuantity: int(bookEntity.StockQuantity),
		PurchaseCount: int(bookEntity.PurchaseCount),
		Rating:        float64(bookEntity.Rating),
		CreatedAt:     bookEntity.CreatedAt.Time,
		UpdatedAt:     bookEntity.UpdatedAt.Time,
		DeletedAt:     bookEntity.DeletedAt.Time,
	}
}

func (r *Book) fillBooksCategories(
	ctx context.Context,
	books *[]domain.Book,
) error {
	bookIDs := make([]uuid.UUID, 0, len(*books))
	for _, book := range *books {
		bookIDs = append(bookIDs, book.ID)
	}
	booksCategoriesMap := make(map[uuid.UUID][]uuid.UUID)
	bookCategoryEntities, err := r.queries.ListBooksCategories(ctx, sqlc.ListBooksCategoriesParams{
		BookIDs: bookIDs,
	})
	if err != nil {
		return err
	}
	for _, bookCategoryEntity := range bookCategoryEntities {
		booksCategoriesMap[bookCategoryEntity.BookID] = append(
			booksCategoriesMap[bookCategoryEntity.BookID],
			bookCategoryEntity.CategoryID,
		)
	}
	for i, book := range *books {
		if categoryIDs, ok := booksCategoriesMap[book.ID]; ok {
			(*books)[i].CategoryIDs = categoryIDs
		}
	}
	return nil
}

func (r *Book) fillBooksMedium(
	ctx context.Context,
	books *[]domain.Book,
) error {
	bookIDs := make([]uuid.UUID, 0, len(*books))
	for _, book := range *books {
		bookIDs = append(bookIDs, book.ID)
	}
	booksMediumMap := make(map[uuid.UUID][]domain.BookMedia)
	bookMediumEntities, err := r.queries.ListBooksMedium(ctx, sqlc.ListBooksMediumParams{
		BookIDs: bookIDs,
	})
	if err != nil {
		return err
	}
	for _, bookMediumEntity := range bookMediumEntities {
		booksMediumMap[bookMediumEntity.BookID] = append(
			booksMediumMap[bookMediumEntity.BookID],
			domain.BookMedia{
				MediaID: bookMediumEntity.MediaID,
				IsCover: bookMediumEntity.IsCover,
				Order:   int(bookMediumEntity.Order),
			},
		)
	}
	for i, book := range *books {
		if medium, ok := booksMediumMap[book.ID]; ok {
			(*books)[i].Medium = medium
		}
	}
	return nil
}

func (r *Book) mergeCategories(
	ctx context.Context,
	qtx *sqlc.Queries,
	book *domain.Book,
) error {
	if err := qtx.CreateTempTableBooksCategories(ctx); err != nil {
		return err
	}
	param := make([]sqlc.InsertTempTableBooksCategoriesParams, 0, len(book.CategoryIDs))
	for _, categoryID := range book.CategoryIDs {
		param = append(param, sqlc.InsertTempTableBooksCategoriesParams{
			BookID:     book.ID,
			CategoryID: categoryID,
		})
	}
	if _, err := qtx.InsertTempTableBooksCategories(ctx, param); err != nil {
		return err
	}
	if err := qtx.MergeBooksCategoriesFromTemp(ctx); err != nil {
		return err
	}
	return nil
}

func (r *Book) mergeMedium(
	ctx context.Context,
	qtx *sqlc.Queries,
	book *domain.Book,
) error {
	if err := qtx.CreateTempTableBooksMedium(ctx); err != nil {
		return err
	}
	param := make([]sqlc.InsertTempTableBooksMediumParams, 0, len(book.Medium))
	for _, medium := range book.Medium {
		param = append(param, sqlc.InsertTempTableBooksMediumParams{
			BookID:  book.ID,
			MediaID: medium.MediaID,
			Order:   int32(medium.Order),
			IsCover: medium.IsCover,
		})
	}
	if _, err := qtx.InsertTempTableBooksMedium(ctx, param); err != nil {
		return err
	}
	if err := qtx.MergeBooksMediumFromTemp(ctx); err != nil {
		return err
	}
	return nil
}
