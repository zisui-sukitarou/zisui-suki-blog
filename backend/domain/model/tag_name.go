package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TagName string

func NewTagName(tagName string) (TagName, error) {
	if err := validation.Validate(tagName,
		validation.Required,
		validation.Length(1, 15),
		validation.Match(regexp.MustCompile("^[\u3041-\u3096\u30A1-\u30FA\u3400-\u9FFF\uF900-\uFAFFー々〇〻]+$")),
	); err != nil {
		return TagName(""), err
	}

	return TagName(tagName), nil
}