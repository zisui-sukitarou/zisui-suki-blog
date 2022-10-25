package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserIcon string

func NewUserIcon(icon string) (UserIcon, error) {
	if err := validation.Validate(icon,
		validation.Required,
		is.URL,
	); err != nil {
		return UserIcon(""), err
	}

	return UserIcon(icon), nil
}
