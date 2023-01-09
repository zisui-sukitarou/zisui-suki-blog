package model

import (
	"log"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TagName string

func NewTagName(tagName string) (TagName, error) {
	log.Println("model: tag: length of tagName:", len(tagName))
	if err := validation.Validate(tagName,
		validation.Required,
		validation.Length(1, 45),
		validation.Match(regexp.MustCompile("^[\u3041-\u3096\u30A1-\u30FA\u3400-\u9FFF\uF900-\uFAFFー々〇〻]+$")),
	); err != nil {
		return TagName(""), err
	}

	return TagName(tagName), nil
}