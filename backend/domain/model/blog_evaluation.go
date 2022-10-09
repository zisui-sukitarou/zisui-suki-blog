package model

import validation "github.com/go-ozzo/ozzo-validation"

type BlogEvaluation int

func NewBlogEvaluation(evaluation int) (BlogEvaluation, error) {
	if err := validation.Validate(evaluation,
		validation.Min(0),
		validation.Max(10),
	); err != nil {
		return BlogEvaluation(0), err
	}

	return BlogEvaluation(evaluation), nil
}
