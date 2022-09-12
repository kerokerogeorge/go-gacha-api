package helper

import (
	"math/rand"

	"github.com/Songmu/flextime"
	"github.com/oklog/ulid"
)

func NewULID() ulid.ULID {
	t := flextime.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}
