package model

type DraftStatus uint

func NewDraftStatus(status uint) (DraftStatus, error) {
	return DraftStatus(status), nil
}