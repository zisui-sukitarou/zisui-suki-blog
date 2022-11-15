package model

import "errors"

type DraftEvaluation uint

func NewDraftEvaluation(evaluation uint) (DraftEvaluation, error) {
	if evaluation < 0 {
		return DraftEvaluation(evaluation), errors.New("evaluation must be greater or equal than 0")
	}
	if evaluation > 10 {
		return DraftEvaluation(evaluation), errors.New("evaluation must be smaller or equal than 10")
	}

	return DraftEvaluation(evaluation), nil
}
