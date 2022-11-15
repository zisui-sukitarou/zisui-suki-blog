package model

type DraftTitle string

func NewDraftTitle(title string) (DraftTitle, error) {
	return DraftTitle(title), nil
}
