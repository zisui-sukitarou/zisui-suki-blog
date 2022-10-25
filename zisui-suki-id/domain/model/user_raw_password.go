package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserRawPassword string

func NewUserRawPassword(userRawPassword string) (UserRawPassword, error){
	if err := validation.Validate(userRawPassword,
		validation.Required,
		validation.Length(6, 15),
		validation.Match(regexp.MustCompile("^[a-zA-Z0-9!?_@-]+$")),
	); err != nil {
		return UserRawPassword(""), err
	}

	return UserRawPassword(userRawPassword), nil
}