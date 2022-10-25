package model

import "errors"

type BlogEvaluation uint

func NewBlogEvaluation(evaluation uint) (BlogEvaluation, error) {
	if evaluation < 0 {
		return BlogEvaluation(evaluation), errors.New("evaluation must be greater or equal than 0")
	}
	if evaluation > 10 {
		return BlogEvaluation(evaluation), errors.New("evaluation must be smaller or equal than 10")
	}

	return BlogEvaluation(evaluation), nil
}
