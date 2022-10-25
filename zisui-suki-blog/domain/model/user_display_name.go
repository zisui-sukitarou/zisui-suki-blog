package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserDisplayName string 

func NewUserDisplayName(userName string) (UserDisplayName, error){
	if err := validation.Validate(userName,
		validation.Required,
		validation.Length(1, 15),
	); err != nil {
		return UserDisplayName(""), err
	}

	return UserDisplayName(userName), nil
}