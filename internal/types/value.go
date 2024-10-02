package types

import (
	"time"
)

type Value struct {
	Val          []byte
	DateCreated  time.Time
	DateModified time.Time
}
