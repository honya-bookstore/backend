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
	DeletedAt *time.Time   `json:"deletedAt" validate:"omitnil,gtefield=CreatedAt"`
}

type ReviewVote struct {
	ID        uuid.UUID `json:"id"        binding:"required" validate:"required"`
	UserID    uuid.UUID `json:"userId"    binding:"required" validate:"required,uuid"`
	IsUp      bool      `json:"isUp"      binding:"required" validate:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required" validate:"required"`
}
