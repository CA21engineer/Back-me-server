package util

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func GenerateUlid() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return id.String()
}
