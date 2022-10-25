package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserName string

func NewUserName(name string) (UserName, error) {
	if err := validation.Validate(name,
		validation.Required,
		validation.Length(6, 15),
		validation.Match(regexp.MustCompile("^[a-zA-Z0-9_-]+$")),
	); err != nil{
		return UserName(""), err
	}

	return UserName(name), nil
}