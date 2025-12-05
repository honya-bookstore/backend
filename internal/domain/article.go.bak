package domain

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID          uuid.UUID `validate:"required"`
	Slug        string    `validate:"required,gte=2,lte=200"`
	Title       string    `validate:"required,gte=2,lte=200"`
	ThumbnailID uuid.UUID `validate:"required"`
	Content     string    `validate:"omitempty,lte=50000"`
	Tags        []string  `validate:"omitempty,dive,gte=1,lte=50"`
	UserID      uuid.UUID `validate:"required,uuid"`
	CreatedAt   time.Time `validate:"required"`
	UpdatedAt   time.Time `validate:"required,gtefield=CreatedAt"`
	DeletedAt   time.Time `validate:"omitempty,gtefield=CreatedAt"`
}

func NewArticle(
	slug string,
	title string,
	thumbnailID uuid.UUID,
	content string,
	tags []string,
	userID uuid.UUID,
) (*Article, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &Article{
		ID:          id,
		Slug:        slug,
		Title:       title,
		ThumbnailID: thumbnailID,
		Content:     content,
		Tags:        tags,

		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (a *Article) Update(
	slug string,
	title string,
	thumbnailID uuid.UUID,
	content string,
) {
	updated := false
	if slug != "" && slug != a.Slug {
		a.Slug = slug
		updated = true
	}
	if title != "" && title != a.Title {
		a.Title = title
		updated = true
	}
	if thumbnailID != uuid.Nil && thumbnailID != a.ThumbnailID {
		a.ThumbnailID = thumbnailID
		updated = true
	}
	if content != "" && content != a.Content {
		a.Content = content
		updated = true
	}

	if updated {
		a.UpdatedAt = time.Now()
	}
}

func (a *Article) AddTags(newTags ...string) {
	a.Tags = append(a.Tags, newTags...)
}

func (a *Article) RemoveTags(tagsToRemove ...string) {
	tagMap := make(map[string]struct{})
	for _, tag := range tagsToRemove {
		tagMap[tag] = struct{}{}
	}

	filteredTags := make([]string, 0, len(a.Tags)-len(tagsToRemove))
	for _, tag := range a.Tags {
		if _, found := tagMap[tag]; !found {
			filteredTags = append(filteredTags, tag)
		}
	}
	a.Tags = filteredTags
}

func (a *Article) Remove() {
	a.DeletedAt = time.Now()
}
