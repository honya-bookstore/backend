package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Category struct {
	queries *sqlc.Queries
	conn    *pgxpool.Pool
}

var _ domain.CategoryRepository = (*Category)(nil)

func ProvideCategory(
	queries *sqlc.Queries,
	conn *pgxpool.Pool,
) *Category {
	return &Category{
		queries: queries,
		conn:    conn,
	}
}

func (r *Category) List(
	ctx context.Context,
	params domain.CategoryRepositoryListParam,
) (*[]domain.Category, error) {
	categoryEntitities, err := r.queries.ListCategories(ctx, sqlc.ListCategoriesParams{
		IDs:    params.CategoryIDs,
		Search: params.Search,
		Limit:  int32(params.Limit),
		Offset: int32(params.Offset),
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	categories := make([]domain.Category, 0, len(categoryEntitities))
	for _, categoryEntity := range categoryEntitities {
		categories = append(categories, *toCategoryDomain(&categoryEntity))
	}
	return &categories, nil
}

func (r *Category) Count(
	ctx context.Context,
) (*int, error) {
	count, err := r.queries.CountCategories(ctx, sqlc.CountCategoriesParams{})
	if err != nil {
		return nil, toDomainError(err)
	}
	countInt := int(count)
	return &countInt, nil
}

func (r *Category) Get(
	ctx context.Context,
	params domain.CategoryRepositoryGetParam,
) (*domain.Category, error) {
	categoryEntity, err := r.queries.GetCategory(ctx, sqlc.GetCategoryParams{
		ID:   params.CategoryID,
		Slug: params.CategorySlug,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	category := toCategoryDomain(&categoryEntity)
	return category, nil
}

func (r *Category) Save(
	ctx context.Context,
	params domain.CategoryRepositorySaveParam,
) error {
	err := r.queries.UpsertCategory(
		ctx, sqlc.UpsertCategoryParams{
			ID:          params.Category.ID,
			Name:        params.Category.Name,
			Slug:        params.Category.Slug,
			Description: ptr.To(params.Category.Description),
			CreatedAt: pgtype.Timestamptz{
				Time:  params.Category.CreatedAt,
				Valid: true,
			},
			UpdatedAt: pgtype.Timestamptz{
				Time:  params.Category.UpdatedAt,
				Valid: true,
			},
			DeletedAt: pgtype.Timestamptz{
				Time:  params.Category.DeletedAt,
				Valid: !params.Category.DeletedAt.IsZero(),
			},
		},
	)
	if err != nil {
		return toDomainError(err)
	}
	return nil
}

func toCategoryDomain(categoryEntity *sqlc.Category) *domain.Category {
	return &domain.Category{
		ID:          categoryEntity.ID,
		Name:        categoryEntity.Name,
		Slug:        categoryEntity.Slug,
		Description: ptr.Deref(categoryEntity.Description, ""),
		CreatedAt:   categoryEntity.CreatedAt.Time,
		UpdatedAt:   categoryEntity.UpdatedAt.Time,
		DeletedAt:   categoryEntity.DeletedAt.Time,
	}
}
