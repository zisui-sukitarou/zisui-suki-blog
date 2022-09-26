package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserId string

func NewUserId(userId string) (UserId, error) {
	if err := validation.Validate(userId,
		validation.Required,
		validation.Length(6, 15),
		validation.Match(regexp.MustCompile("^[a-zA-Z0-9_-]+$")),
	); err != nil{
		return UserId(""), err
	}

	return UserId(userId), nil
}
