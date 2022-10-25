package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type BlogContent string

func NewBlogContent(content string) (BlogContent, error) {
	if err := validation.Validate(content,
		validation.Required,
	); err != nil {
		return BlogContent(""), err
	}

	return BlogContent(content), nil
}
