package model

type DraftContent string

func NewDraftContent(content string) (DraftContent, error) {
	return DraftContent(content), nil
}
