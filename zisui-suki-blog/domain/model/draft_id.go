package model

type DraftId string

func NewDraftId(id string) (DraftId, error) {
	return DraftId(id), nil
}
