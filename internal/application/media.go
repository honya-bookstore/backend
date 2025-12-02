package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type Media struct {
	mediaRepo    domain.MediaRepository
	mediaService domain.MediaService
}

func ProvideMedia(
	mediaRepo domain.MediaRepository,
	mediaService domain.MediaService,
) *Media {
	return &Media{
		mediaRepo:    mediaRepo,
		mediaService: mediaService,
	}
}

var _ http.MediaApplication = &Media{}

func (m *Media) List(ctx context.Context, param http.ListMediaRequestDTO) (*http.PaginationResponseDTO[http.MediaResponseDTO], error) {
	queryParams := param.QueryParams
	if queryParams == nil {
		queryParams = &http.ListMediaRequestQueryParams{}
	}

	media, err := m.mediaRepo.List(
		ctx,
		domain.MediaRepositoryListParam{
			Search: queryParams.Search,
			Limit:  queryParams.Limit,
			Offset: (queryParams.Page - 1) * queryParams.Limit,
		},
	)
	if err != nil {
		return nil, err
	}

	count, err := m.mediaRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	mediaDtos := http.ToMediaResponseDTOList(*media)

	pagination := newPaginationResponseDto(
		mediaDtos,
		*count,
		queryParams.Page,
		queryParams.Limit,
	)

	return pagination, nil
}

func (m *Media) Get(ctx context.Context, param http.GetMediaRequestDTO) (*http.MediaResponseDTO, error) {
	mediaID, err := uuid.Parse(param.PathParams.MediaID)
	if err != nil {
		return nil, domain.ErrInvalid
	}

	media, err := m.mediaRepo.Get(ctx, domain.MediaRepositoryGetParam{
		MediaID: mediaID,
	})
	if err != nil {
		return nil, err
	}

	return http.ToMediaResponseDTO(media), nil
}

func (m *Media) Create(ctx context.Context, param http.CreateMediaRequestDTO) (*http.MediaResponseDTO, error) {
	media, err := domain.NewMedia(
		param.Data.URL,
		param.Data.AltText,
		param.Data.Order,
	)
	if err != nil {
		return nil, err
	}

	if err := m.mediaService.Validate(*media); err != nil {
		return nil, err
	}

	err = m.mediaRepo.Save(ctx, domain.MediaRepositorySaveParam{
		Media: *media,
	})
	if err != nil {
		return nil, err
	}

	return http.ToMediaResponseDTO(media), nil
}

func (m *Media) Delete(ctx context.Context, param http.DeleteMediaRequestDTO) error {
	mediaID, err := uuid.Parse(param.PathParams.MediaID)
	if err != nil {
		return domain.ErrInvalid
	}

	media, err := m.mediaRepo.Get(ctx, domain.MediaRepositoryGetParam{
		MediaID: mediaID,
	})
	if err != nil {
		return err
	}

	media.Delete()

	if err := m.mediaService.Validate(*media); err != nil {
		return err
	}

	err = m.mediaRepo.Save(ctx, domain.MediaRepositorySaveParam{
		Media: *media,
	})
	if err != nil {
		return err
	}

	return nil
}
