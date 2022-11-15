package model

type DraftAbstract string

func NewDraftAbstract(abstract string) (DraftAbstract, error) {
	return DraftAbstract(abstract), nil
}
