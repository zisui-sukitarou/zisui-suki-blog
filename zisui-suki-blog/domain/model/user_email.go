package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserEmail string 

func NewUserEmail(userEmail string) (UserEmail, error){
	if err := validation.Validate(userEmail,
		validation.Required,
		is.Email,
	); err != nil{
		return UserEmail(""), err
	}

	return UserEmail(userEmail), nil
}