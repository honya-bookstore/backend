package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"
)

type Category struct {
	categoryRepo    domain.CategoryRepository
	categoryService domain.CategoryService
}

func ProvideCategory(
	categoryRepo domain.CategoryRepository,
	categoryService domain.CategoryService,
) *Category {
	return &Category{
		categoryRepo:    categoryRepo,
		categoryService: categoryService,
	}
}

var _ http.CategoryApplication = &Category{}

func (c *Category) Create(ctx context.Context, param http.CreateCategoryRequestDTO) (*http.CategoryResponseDTO, error) {
	category, err := domain.NewCategory(
		param.Data.Slug,
		param.Data.Name,
		param.Data.Description,
	)
	if err != nil {
		return nil, err
	}

	if err := c.categoryService.Validate(*category); err != nil {
		return nil, err
	}

	err = c.categoryRepo.Save(ctx, domain.CategoryRepositorySaveParam{
		Category: *category,
	})
	if err != nil {
		return nil, err
	}

	return http.ToCategoryResponseDTO(category), nil
}

func (c *Category) List(ctx context.Context, param http.ListCategoryRequestDTO) (*http.PaginationResponseDTO[http.CategoryResponseDTO], error) {
	queryParams := param.QueryParams
	if queryParams == nil {
		queryParams = &http.ListCategoryRequestQueryParams{}
	}

	categories, err := c.categoryRepo.List(
		ctx,
		domain.CategoryRepositoryListParam{
			Search: queryParams.Search,
			Limit:  queryParams.Limit,
			Offset: (queryParams.Page - 1) * queryParams.Limit,
		},
	)
	if err != nil {
		return nil, err
	}

	count, err := c.categoryRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	categoryDtos := http.ToCategoryResponseDTOList(*categories)

	pagination := newPaginationResponseDto(
		categoryDtos,
		*count,
		queryParams.Page,
		queryParams.Limit,
	)

	return pagination, nil
}

func (c *Category) GetBySlug(ctx context.Context, param http.GetCategoryRequestDTO) (*http.CategoryResponseDTO, error) {
	category, err := c.categoryRepo.Get(ctx, domain.CategoryRepositoryGetParam{
		CategorySlug: param.PathParams.Slug,
	})
	if err != nil {
		return nil, err
	}

	return http.ToCategoryResponseDTO(category), nil
}

func (c *Category) Update(ctx context.Context, param http.UpdateCategoryRequestDTO) (*http.CategoryResponseDTO, error) {
	category, err := c.categoryRepo.Get(ctx, domain.CategoryRepositoryGetParam{
		CategoryID: param.PathParams.CategoryID,
	})
	if err != nil {
		return nil, err
	}

	category.Update(
		param.Data.Name,
		param.Data.Description,
		param.Data.Slug,
	)

	if err := c.categoryService.Validate(*category); err != nil {
		return nil, err
	}

	err = c.categoryRepo.Save(ctx, domain.CategoryRepositorySaveParam{
		Category: *category,
	})
	if err != nil {
		return nil, err
	}

	return http.ToCategoryResponseDTO(category), nil
}

func (c *Category) Delete(ctx context.Context, param http.DeleteCategoryRequestDTO) error {
	category, err := c.categoryRepo.Get(ctx, domain.CategoryRepositoryGetParam{
		CategoryID: param.PathParams.CategoryID,
	})
	if err != nil {
		return err
	}

	category.Delete()

	if err := c.categoryService.Validate(*category); err != nil {
		return err
	}

	err = c.categoryRepo.Save(ctx, domain.CategoryRepositorySaveParam{
		Category: *category,
	})
	if err != nil {
		return err
	}

	return nil
}
