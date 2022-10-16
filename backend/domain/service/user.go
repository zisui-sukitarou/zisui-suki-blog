package service

import (
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
}

func (u *User) Exists(
	userId model.UserId,
	repository repository.UserRepository,
) (
	bool,
	error,
) {
	exists, _, err := repository.FindById(userId)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (u *User) HashPassword(
	rawPassword model.UserRawPassword,
) (
	model.UserHashedPassword,
	error,
) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return model.UserHashedPassword(""), err
	}

	return model.UserHashedPassword(hashedPassword), nil
}

func (u *User) IsConfigured(
	hashedPassword model.UserHashedPassword,
	inputPassword model.UserRawPassword,
) (
	bool,
	error,
) {
	/* hash(input) == hashed_password ? */
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(inputPassword),
	); err != nil {
		return false, err
	}

	return true, nil
}
