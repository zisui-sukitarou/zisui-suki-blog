package service

import (
	"math/rand"
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"

	"github.com/oklog/ulid/v2"
)

type User struct {
}

func (u *User) GenULID() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return ulid.New(ms, entropy)
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