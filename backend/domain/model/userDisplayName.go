package model

import validation "github.com/go-ozzo/ozzo-validation"

type UserDisplayName string

func NewUserDisplayName(userDisplayName string) (UserDisplayName, error) {
	if err := validation.Validate(userDisplayName,
		validation.Required,
		validation.Length(1, 15),
	); err != nil {
		return UserDisplayName(""), err
	}
	
	return UserDisplayName(userDisplayName), nil
}