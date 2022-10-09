package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type BlogTitle string

func NewBlogTitle(title string) (BlogTitle, error) {
	if err := validation.Validate(title,
		validation.Required,
	); err != nil {
		return BlogTitle(""), err
	}

	return BlogTitle(title), nil
}
