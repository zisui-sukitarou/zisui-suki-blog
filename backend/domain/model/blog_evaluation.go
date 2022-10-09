package model

import validation "github.com/go-ozzo/ozzo-validation"

type BlogEvaluation uint

func NewBlogEvaluation(evaluation uint) (BlogEvaluation, error) {
	if err := validation.Validate(evaluation,
		validation.Min(0),
		validation.Max(10),
	); err != nil {
		return BlogEvaluation(0), err
	}

	return BlogEvaluation(evaluation), nil
}
