package domain

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID    `json:"id"        binding:"required"                    validate:"required"`
	Rating    int16        `json:"rating"    binding:"required"                    validate:"required,gte=1,lte=5"`
	VoteCount int          `json:"voteCount" binding:"required"                    validate:"required,gte=0"`
	Content   string       `json:"content"   validate:"omitempty,lte=2000"`
	UserID    uuid.UUID    `json:"userId"    binding:"required"                    validate:"required,uuid"`
	BookID    uuid.UUID    `json:"bookId"    binding:"required"                    validate:"required,uuid"`
	Votes     []ReviewVote `json:"votes"     validate:"omitempty,dive"`
	CreatedAt time.Time    `json:"createdAt" binding:"required"                    validate:"required"`
	UpdatedAt time.Time    `json:"updatedAt" binding:"required"                    validate:"required,gtefield=CreatedAt"`
	DeletedAt time.Time    `json:"deletedAt" validate:"omitnil,gtefield=CreatedAt"`
}

type ReviewVote struct {
	ID        uuid.UUID `json:"id"        binding:"required" validate:"required"`
	UserID    uuid.UUID `json:"userId"    binding:"required" validate:"required,uuid"`
	IsUp      bool      `json:"isUp"      binding:"required" validate:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required" validate:"required"`
}

func NewReview(
	rating int16,
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

func (r *Review) Update(rating *int16, content *string) {
	updated := false
	if rating != nil {
		r.Rating = *rating
		updated = true
	}
	if content != nil {
		r.Content = *content
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
