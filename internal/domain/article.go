package domain

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID          uuid.UUID  `json:"id"          binding:"required"                     validate:"required"`
	Slug        string     `json:"slug"        binding:"required"                     validate:"required,gte=2,lte=200"`
	Title       string     `json:"title"       binding:"required"                     validate:"required,gte=2,lte=200"`
	ThumbnailID uuid.UUID  `json:"thumbnailId" validate:"required"`
	Content     string     `json:"content"     validate:"omitempty,lte=50000"`
	Tags        []string   `json:"tags"        validate:"omitempty,dive,gte=1,lte=50"`
	UserID      uuid.UUID  `json:"userId"      binding:"required"                     validate:"required,uuid"`
	CreatedAt   time.Time  `json:"createdAt"   binding:"required"                     validate:"required"`
	UpdatedAt   time.Time  `json:"updatedAt"   binding:"required"                     validate:"required,gtefield=CreatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"   validate:"omitnil,gtefield=CreatedAt"`
}

func NewArticle(slug string, title string, thumbnailID uuid.UUID, content string, tags []string, userID uuid.UUID) (*Article, error) {
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
		UserID:      userID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (a *Article) Update(slug *string, title *string, thumbnailID *uuid.UUID, content *string) {
	updated := false
	if slug != nil {
		a.Slug = *slug
		updated = true
	}
	if title != nil {
		a.Title = *title
		updated = true
	}
	if thumbnailID != nil {
		a.ThumbnailID = *thumbnailID
		updated = true
	}
	if content != nil {
		a.Content = *content
		updated = true
	}

	if updated {
		a.UpdatedAt = time.Now()
	}
}

func (a *Article) Remove() {
	now := time.Now()
	if a.DeletedAt == nil {
		a.DeletedAt = &now
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
