package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserPassword string

func NewUserPassword(userPassword string) (UserPassword, error){
	if err := validation.Validate(userPassword,
		validation.Required,
		validation.Length(6, 15),
		validation.Match(regexp.MustCompile("^[a-zA-Z0-9!?_@-]+$")),
	); err != nil {
		return UserPassword(""), err
	}

	return UserPassword(userPassword), nil
}