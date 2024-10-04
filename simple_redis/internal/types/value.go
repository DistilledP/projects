package types

import (
	"time"
)

type Value struct {
	Val          []byte
	DateCreated  time.Time
	DateModified time.Time
}

func (v *Value) Update(val string) {
	v.Val = []byte(val)
	v.DateModified = time.Now()
}

func NewValue(val string) Value {
	now := time.Now()
	return Value{
		Val:          []byte(val),
		DateCreated:  now,
		DateModified: now,
	}
}
