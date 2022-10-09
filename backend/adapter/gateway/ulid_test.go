package gateway_test

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

func genUlid() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return ulid.New(ms, entropy)
}

func TestGenUlid(t *testing.T) {
	for i := 0; i < 5; i++ {
		id, err := genUlid()
		if err != nil {
			t.Errorf(err.Error())
		}
		log.Println(id.String())
		time.Sleep(time.Millisecond * 200)
	}
}