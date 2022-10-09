package service

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type Blog struct{}

func (b *Blog) GenULID() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return ulid.New(ms, entropy)
}
