package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserName string 

func NewUserName(userName string) (UserName, error){
	if err := validation.Validate(userName,
		validation.Required,
		validation.Length(1, 15),
	); err != nil {
		return UserName(""), err
	}

	return UserName(userName), nil
}