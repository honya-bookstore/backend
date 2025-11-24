package domain

type DeletedParam string

const (
	DeletedExcludeParam DeletedParam = "exclude"
	DeletedOnlyParam    DeletedParam = "only"
	DeletedAllParam     DeletedParam = "all"
)
