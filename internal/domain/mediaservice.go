package domain

type MediaService interface {
	Validate(
		media Media,
	) error
}
