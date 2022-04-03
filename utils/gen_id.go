package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"github.com/rs/xid"
)

func GenUlid() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

func GenXid() string {
	id := xid.New()
	return id.String()
}
