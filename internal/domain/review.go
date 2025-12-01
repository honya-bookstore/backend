package domain

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID    `validate:"required"`
	Rating    uint8        `validate:"required,gte=1,lte=5"`
	VoteCount int          `validate:"required,gte=0"`
	Content   string       `validate:"omitempty,lte=2000"`
	UserID    uuid.UUID    `validate:"required"`
	BookID    uuid.UUID    `validate:"required"`
	Votes     []ReviewVote `validate:"omitempty,dive"`
	CreatedAt time.Time    `validate:"required"`
	UpdatedAt time.Time    `validate:"required,gtefield=CreatedAt"`
	DeletedAt time.Time    `validate:"omitempty,gtefield=CreatedAt"`
}

type ReviewVote struct {
	ID        uuid.UUID `validate:"required"`
	UserID    uuid.UUID `validate:"required,uuid"`
	IsUp      bool      `validate:"required"`
	CreatedAt time.Time `validate:"required"`
}

func NewReview(
	rating uint8,
	content string,
	userID uuid.UUID,
	bookID uuid.UUID,
) (*Review, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &Review{
		ID:        id,
		Rating:    rating,
		VoteCount: 0,
		Content:   content,
		UserID:    userID,
		BookID:    bookID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func NewReviewVote(userID uuid.UUID, isUp bool) (*ReviewVote, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &ReviewVote{
		ID:        id,
		UserID:    userID,
		IsUp:      isUp,
		CreatedAt: now,
	}, nil
}

func (r *Review) Update(rating uint8, content string) {
	updated := false
	if r.Rating != rating && rating != 0 {
		r.Rating = rating
		updated = true
	}
	if r.Content != content && content != "" {
		r.Content = content
		updated = true
	}

	if updated {
		r.UpdatedAt = time.Now()
	}
}

func (r *Review) AddVote(vote ReviewVote) {
	r.Votes = append(r.Votes, vote)
	if vote.IsUp {
		r.VoteCount++
		return
	}
	r.VoteCount--
}

func (r *ReviewVote) Update(isUp bool) {
	r.IsUp = isUp
}

func (r *Review) Remove() {
	now := time.Now()
	r.DeletedAt = now
	r.UpdatedAt = now
}
