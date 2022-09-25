package model

import validation "github.com/go-ozzo/ozzo-validation"

type TagName string

func NewTagName(tagName string) (TagName, error) {
	if err := validation.Validate(tagName,
		validation.Required,
		validation.Length(1, 15),
	); err != nil {
		return TagName(""), err
	}

	return TagName(tagName), nil
}