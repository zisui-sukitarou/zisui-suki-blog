package service

import (
	"math/rand"
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
}

type Status struct {
}

func (u *User) GenULID() (model.UserId, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	if err != nil {
		return model.UserId(""), err
	}

	return model.NewUserId(id.String())
}

func (u *User) Exists(
	name model.UserName,
	repository repository.UserRepository,
) (
	bool,
	error,
) {
	exists, _, err := repository.FindByName(name)
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

// TODO
func (u *User) GenJWT(userId model.UserId) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("SECRET_KEY"))
}
