package service

import (
	"math/rand"
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"

	"github.com/oklog/ulid/v2"
)

type Draft struct{}

func (d *Draft) GenULID() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return ulid.New(ms, entropy)
}

func (d *Draft) Exists(blogId model.BlogId, repository repository.BlogRepository) (bool, error) {
	exists, _, err := repository.FindById(blogId)
	if err != nil {
		return false, err
	}
	return exists, nil
}
