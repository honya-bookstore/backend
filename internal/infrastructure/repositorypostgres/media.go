package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Media struct {
	queries *sqlc.Queries
	conn    *pgxpool.Pool
}

var _ domain.MediaRepository = (*Media)(nil)

func ProvideMedia(
	queries *sqlc.Queries,
	conn *pgxpool.Pool,
) *Media {
	return &Media{
		queries: queries,
		conn:    conn,
	}
}

func (r *Media) List(
	ctx context.Context,
	params domain.MediaRepositoryListParam,
) (*[]domain.Media, error) {
	mediaEntities, err := r.queries.ListMedium(ctx, sqlc.ListMediumParams{
		IDs:    params.MediaIDs,
		Offset: int32(params.Offset),
		Limit:  int32(params.Limit),
		Search: params.Search,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	medias := make([]domain.Media, 0, len(mediaEntities))
	for _, mediaEntity := range mediaEntities {
		medias = append(medias, *toMediaDomain(&mediaEntity))
	}
	return &medias, nil
}

func (r *Media) Count(
	ctx context.Context,
) (*int, error) {
	count, err := r.queries.CountMedium(ctx, sqlc.CountMediumParams{})
	if err != nil {
		return nil, toDomainError(err)
	}
	countInt := int(count)
	return &countInt, nil
}

func (r *Media) Get(
	ctx context.Context,
	params domain.MediaRepositoryGetParam,
) (*domain.Media, error) {
	mediaEntity, err := r.queries.GetMedia(ctx, sqlc.GetMediaParams{
		ID: params.MediaID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	media := toMediaDomain(&mediaEntity)
	return media, nil
}

func (r *Media) Save(
	ctx context.Context,
	params domain.MediaRepositorySaveParam,
) error {
	err := r.queries.UpsertMedia(ctx, sqlc.UpsertMediaParams{
		ID:      params.Media.ID,
		URL:     params.Media.URL,
		AltText: params.Media.AltText,
		CreatedAt: pgtype.Timestamptz{
			Time:  params.Media.CreatedAt,
			Valid: true,
		},
		DeletedAt: pgtype.Timestamptz{
			Time:  params.Media.DeletedAt,
			Valid: true,
		},
	},
	)
	if err != nil {
		return toDomainError(err)
	}
	return nil
}

func toMediaDomain(mediaEntity *sqlc.Medium) *domain.Media {
	return &domain.Media{
		ID:        mediaEntity.ID,
		URL:       mediaEntity.URL,
		AltText:   ptr.Deref(mediaEntity.AltText, ""),
		CreatedAt: mediaEntity.CreatedAt.Time,
		DeletedAt: mediaEntity.DeletedAt.Time,
	}
}
