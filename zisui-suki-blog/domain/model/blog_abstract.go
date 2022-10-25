package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type BlogAbstract string

func NewBlogAbstract(abstract string) (BlogAbstract, error) {

	if err := validation.Validate(abstract,
		validation.Required,
	); err != nil {
		return BlogAbstract(""), err
	}

	return BlogAbstract(abstract), nil
}
